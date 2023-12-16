package handlers

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/netip"
	"server/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getIPAddress(r *http.Request) (*netip.Addr, error) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return nil, err
	}

	return &addr, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Login logs the user in by creating a session
func (h *Handler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Queries.GetUserByUsername(context.Background(), input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	parsedIpAddress, err := getIPAddress(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IP address"})
		return
	}

	sessionID, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	sessionIDBytes, err := sessionID.MarshalBinary()
	if err != nil {
		log.Fatalf("failed to convert UUID to byte array: %v", err)
	}

	var sessionID16Bytes [16]byte
	copy(sessionID16Bytes[:], sessionIDBytes[:])

	createSessionParams := db.CreateUserSessionParams{
		SessionID:  pgtype.UUID{Bytes: sessionID16Bytes, Valid: true},
		UserID:     pgtype.Int4{Int32: user.UserID, Valid: true},
		ExpiryDate: pgtype.Timestamptz{Time: time.Now().Add(24 * time.Hour), Valid: true},
		IpAddress:  parsedIpAddress,
		UserAgent:  pgtype.Text{String: c.Request.UserAgent(), Valid: true},
	}

	_, err = h.Queries.CreateUserSession(context.Background(), createSessionParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create session"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID.String(),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	// update last login date
	err = h.Queries.UpdateLastLogin(context.Background(), user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update last login date"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in!"})
}

// Logout logs the user out by invalidating the session
func (h *Handler) Logout(c *gin.Context) {
	// Get session ID from cookie
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No session cookie found"})
		return
	}

	// Convert session ID to UUID
	uuid, err := uuid.Parse(sessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	// Convert UUID to byte array
	uuidBytes, err := uuid.MarshalBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert UUID to byte array"})
		return
	}

	var uuid16Bytes [16]byte
	copy(uuid16Bytes[:], uuidBytes[:])

	// Delete session from database
	err = h.Queries.InvalidateUserSession(context.Background(), pgtype.UUID{Bytes: uuid16Bytes, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invalidate session"})
		return
	}

	// Delete session cookie
	c.SetCookie("session_id", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// InjectRoleNameAndUserID is a middleware that injects the user's role name and user ID into the context
func (h *Handler) InjectRoleNameAndUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.Set("RoleName", "Guest")
			c.Set("Error", fmt.Errorf("no session cookie found"))
			c.Next()
			return
		}

		uuid, err := uuid.Parse(sessionID)
		if err != nil {
			c.Set("RoleName", "Guest")
			c.Set("Error", fmt.Errorf("invalid session ID"))
			c.Next()
			return
		}

		var sessionID16Bytes [16]byte
		copy(sessionID16Bytes[:], uuid[:])

		userSession, err := h.Queries.GetUserSessionAndRoleName(context.Background(), pgtype.UUID{Bytes: sessionID16Bytes, Valid: true})
		if err != nil {
			c.Set("RoleName", "Guest")
			c.Set("Error", fmt.Errorf("invalid session ID"))
			c.Next()
			return
		}

		if userSession.ExpiryDate.Time.Before(time.Now()) {
			c.Set("RoleName", "Guest")
			c.Set("Error", fmt.Errorf("session expired"))
			c.Next()
			return
		}

		c.Set("RoleName", userSession.RoleName)   // type string
		c.Set("UserID", userSession.UserID.Int32) // type int32
		c.Next()
	}
}

// EnsureRole is a middleware that ensures the user has the required role to perform an action
func (h *Handler) EnsureRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("RoleName")
		if userRole == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
			c.Abort()
			return
		}

		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		c.Abort()
	}
}

package handlers

import (
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

func (h *Handler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Queries.GetUserByUsername(*h.Ctx, input.Username)
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

	_, err = h.Queries.CreateUserSession(*h.Ctx, createSessionParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create session"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID.String(),
		HttpOnly: true,
		Secure:   true, // set this to true if you are using https
		Path:     "/",
	})

	// update last login date
	err = h.Queries.UpdateLastLogin(*h.Ctx, user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update last login date"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in!"})
}

func (h *Handler) AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		log.Println(sessionID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, no session_id cookie"})
			c.Abort()
			return
		}

		// Convert sessionID to UUID
		uuid, err := uuid.Parse(sessionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
			c.Abort()
			return
		}

		// Convert UUID to byte array
		var sessionID16Bytes [16]byte
		copy(sessionID16Bytes[:], uuid[:])

		userSession, err := h.Queries.GetUserSession(*h.Ctx, pgtype.UUID{Bytes: sessionID16Bytes, Valid: true})
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, invalid session_id cookie"})
			c.Abort()
			return
		}

		// Check if session has expired
		if userSession.ExpiryDate.Time.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, session has expired"})
			c.Abort()
			return
		}

		user, err := h.Queries.GetUser(*h.Ctx, userSession.UserID.Int32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user role"})
			c.Abort()
			return
		}

		// Inject role_id into gin context
		c.Set("role_id", user.RoleID)

		c.Next()
	}
}

func (h *Handler) UAC(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// COMMENT OUT LATER
		c.Next()

		// role := c.MustGet("role_id").(int32)

		// permissions, err := h.Queries.FetchPermissionNamesForRole(*h.Ctx, role)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	c.Abort()
		// 	return
		// }

		// contains := func(slice []string, item string) bool {
		// 	for _, s := range slice {
		// 		if s == item {
		// 			return true
		// 		}
		// 	}
		// 	return false
		// }

		// if !contains(permissions, requiredPermission) {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		// 	c.Abort()
		// 	return
		// }

		// c.Next()
	}
}

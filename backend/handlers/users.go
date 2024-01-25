package handlers

import (
	"context"
	"net/http"
	"server/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserAPIInput struct {
	Username       string
	Email          string
	Password       string
	ProfilePicture string
	Biography      string
}

func (h *Handler) CreateUser(c *gin.Context) {
	var inputAPI CreateUserAPIInput
	if err := c.ShouldBindJSON(&inputAPI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(inputAPI.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	inputDB := db.CreateUserParams{
		Username:       inputAPI.Username,
		Email:          inputAPI.Email,
		PasswordHash:   hashedPassword,
		ProfilePicture: pgtype.Text{String: inputAPI.ProfilePicture, Valid: true},
		Biography:      pgtype.Text{String: inputAPI.Biography, Valid: true},
		RoleID:         pgtype.Int4{Int32: 3, Valid: true}, // RoleID set directly here to be guest
	}

	err = h.Queries.CreateUser(context.Background(), inputDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	userIDInt, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := h.Queries.GetUser(context.Background(), int32(userIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.Queries.GetAllUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

type UpdateUserExcludingSensitiveAPIParams struct {
	UserID         int32
	Username       string
	Email          string
	ProfilePicture string
	Biography      string
}

func (h *Handler) UpdateUserExcludingSensitive(c *gin.Context) {
	var input UpdateUserExcludingSensitiveAPIParams
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.validateUserID(c, input.UserID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	dbInput := db.UpdateUserExcludingSensitiveParams{
		UserID:         input.UserID,
		Username:       input.Username,
		Email:          input.Email,
		ProfilePicture: pgtype.Text{String: input.ProfilePicture, Valid: true},
		Biography:      pgtype.Text{String: input.Biography, Valid: true},
	}
	err := h.Queries.UpdateUserExcludingSensitive(context.Background(), dbInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

type UpdateUserPasswordAPIParams struct {
	UserID   int32
	Password string
}

func (h *Handler) UpdateUserPassword(c *gin.Context) {
	var input UpdateUserPasswordAPIParams
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.validateUserID(c, input.UserID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	input.Password = hashedPassword

	dbInput := db.UpdateUserPasswordParams{
		UserID:       input.UserID,
		PasswordHash: input.Password,
	}

	err = h.Queries.UpdateUserPassword(context.Background(), dbInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User password updated"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if !h.validateUserID(c, int32(userID)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.Queries.DeleteUser(context.Background(), int32(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

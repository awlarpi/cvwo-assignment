package handlers

import (
	"net/http"
	"server/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserAPIInput struct {
	Username       string
	Email          string
	Password       string
	ProfilePicture string
	Biography      string
	RoleID         int
}

// TODO: RoleID should not be set by the user
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
		RoleID:         pgtype.Int4{Int32: int32(inputAPI.RoleID), Valid: true},
	}

	err = h.Queries.CreateUser(*h.Ctx, inputDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (h *Handler) GetUser(c *gin.Context) {
	var userID int32
	if err := c.ShouldBindUri(&userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.Queries.GetUser(*h.Ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.Queries.GetAllUsers(*h.Ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// update user excluding password
func (h *Handler) UpdateUserExcludingSensitive(c *gin.Context) {
	var input db.UpdateUserExcludingSensitiveParams
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Queries.UpdateUserExcludingSensitive(*h.Ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func (h *Handler) UpdateUserPassword(c *gin.Context) {
	var input db.UpdateUserPasswordParams
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(input.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	input.PasswordHash = hashedPassword

	err = h.Queries.UpdateUserPassword(*h.Ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User password updated"})
}

func (h *Handler) UpdateUserRole(c *gin.Context) {
	var input db.UpdateUserRoleParams
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Queries.UpdateUserRole(*h.Ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User role updated"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var userID int32
	if err := c.ShouldBindUri(&userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Queries.DeleteUser(*h.Ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (h *Handler) DeactivateUser(c *gin.Context) {
	var userID int32
	if err := c.ShouldBindUri(&userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Queries.DeactivateUser(*h.Ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deactivated"})
}

func (h *Handler) ActivateUser(c *gin.Context) {
	var userID int32
	if err := c.ShouldBindUri(&userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Queries.ActivateUser(*h.Ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User activated"})
}

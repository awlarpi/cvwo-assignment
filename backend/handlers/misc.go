// handlers/handlers.go
package handlers

import (
	"net/http"
	"server/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Log     *logrus.Logger
	Dbpool  *pgxpool.Pool
	Queries *db.Queries
}

func (h *Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// validateRoleAndUserID checks if the user is the owner of the resource or an admin
func (h *Handler) validateUserID(c *gin.Context, userIDToModify int32) bool {
	userID, ok := c.Get("userID")
	if !ok {
		return false
	}

	userIDInt32, ok := userID.(int32)
	if !ok {
		return false
	}

	if userIDInt32 != userIDToModify {
		return false
	}

	return true
}

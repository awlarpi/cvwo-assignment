// handlers/handlers.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Log *logrus.Logger
}

func (h *Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (h *Handler) Login(c *gin.Context) {
	// Implement your login logic here
	c.JSON(200, gin.H{
		"message": "Logged in!",
	})
}

func (h *Handler) AuthRequired(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "Bearer valid_token" { // replace with your own token validation logic
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		h.Log.Warn("Unauthorized access attempt")
		return
	}
	c.Next()
}

func (h *Handler) Protected(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Protected route accessed!",
	})
}

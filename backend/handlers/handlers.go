// handlers/handlers.go
package handlers

import (
	"context"
	"net/http"
	"server/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Ctx     *context.Context
	Log     *logrus.Logger
	Queries *db.Queries
}

func (h *Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

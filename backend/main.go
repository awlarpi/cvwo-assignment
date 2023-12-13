package main

import (
	"context"
	"os"
	"server/db"
	"server/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	gin.SetMode(gin.ReleaseMode)
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	r := gin.Default()
	queries := db.New(dbpool)

	h := handlers.Handler{
		Ctx:     &ctx,
		Log:     log,
		Queries: queries,
	}

	api := r.Group("/api")
	{
		api.GET("/ping", h.Ping)

		users := api.Group("/users")
		{
			users.POST("/login", h.Login)
			users.POST("/", h.UAC("CreateUser"), h.CreateUser)
			users.GET("/", h.AuthenticateUser(), h.UAC("GetAllUsers"), h.GetAllUsers)
			users.GET("/:id", h.UAC("GetUser"), h.GetUser)
			users.PUT("/:id", h.UAC("UpdateUserExcludingSensitive"), h.UpdateUserExcludingSensitive)
			users.DELETE("/:id", h.UAC("DeleteUser"), h.DeleteUser)
			users.PUT("/:id/deactivate", h.UAC("DeactivateUser"), h.DeactivateUser)
			users.PUT("/:id/activate", h.UAC("ActivateUser"), h.ActivateUser)
			users.PUT("/password", h.UAC("UpdateUserPassword"), h.UpdateUserPassword)
			users.PUT("/role", h.UAC("UpdateUserRole"), h.UpdateUserRole)
		}

		log.Info("Server running on port 8081")
		r.Run(":8081")
	}
}

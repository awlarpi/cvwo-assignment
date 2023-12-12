package main

import (
	"context"
	"os"
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
	// queries := db.New(dbpool)

	h := handlers.Handler{
		Log: log,
	}

	api := r.Group("/api")
	{
		api.GET("/ping", h.Ping)
		api.POST("/login", h.Login)

		protected := api.Group("/protected")
		protected.Use(h.AuthRequired)
		{
			protected.GET("/test", h.Protected)
		}
	}

	log.Info("Server running on port 8081")
	r.Run(":8081")
}

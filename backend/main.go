package main

import (
	"context"
	"os"
	"server/db"
	"server/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	godotenv.Load() // comment out the line below when using docker-compose

	gin.SetMode(gin.ReleaseMode)
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	err = dbpool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8082", "http://localhost:4173"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	queries := db.New(dbpool)

	h := handlers.Handler{
		Log:     log,
		Queries: queries,
		Dbpool:  dbpool,
	}

	api := r.Group("/api", h.InjectRoleNameAndUserID())
	{
		api.GET("/ping", h.Ping)
		api.POST("/login", h.Login)
		api.POST("/logout", h.Logout)

		users := api.Group("/users")
		{
			users.GET("", h.GetAllUsers)
			users.GET("/:id", h.GetUser)
			users.POST("", h.CreateUser)
			users.PATCH("/password", h.EnsureRole("User", "Moderator", "Admin"), h.UpdateUserPassword)
			users.PUT("", h.EnsureRole("User", "Moderator", "Admin"), h.UpdateUserExcludingSensitive)
			users.DELETE("/:id", h.EnsureRole("User", "Moderator", "Admin"), h.DeleteUser)
		}

		posts := api.Group("/posts")
		{
			posts.GET("", h.GetPostsHandler)
			posts.GET("/:id", h.GetPostHandler)
			posts.GET("/user/:userID", h.GetPostsByUserHandler)
			posts.GET("/category/:postCategoryID", h.GetPostsByCategoryHandler)
			posts.POST("", h.EnsureRole("User", "Moderator", "Admin"), h.CreatePostHandler)
			posts.PUT("", h.EnsureRole("User", "Moderator", "Admin"), h.UpdatePostHandler)
			posts.DELETE("/:id", h.EnsureRole("User", "Moderator", "Admin"), h.DeletePostHandler)
		}

		comments := api.Group("/comments")
		{
			comments.GET("/:commentID", h.GetCommentHandler)
			comments.GET("/post/:postID", h.GetCommentsByPostHandler)
			comments.GET("/user/:userID", h.GetCommentsByUserHandler)
			comments.POST("", h.EnsureRole("User", "Moderator", "Admin"), h.CreateCommentHandler)
			comments.PUT("", h.EnsureRole("User", "Moderator", "Admin"), h.UpdateCommentHandler)
			comments.DELETE("/:commentID", h.EnsureRole("User", "Moderator", "Admin"), h.DeleteCommentHandler)
		}

		log.Info("Server running on port 8081")
		r.Run(":8081")
	}
}

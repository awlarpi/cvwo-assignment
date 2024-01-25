package handlers

import (
	"context"
	"net/http"
	"server/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

// GetPostsHandler handles GET requests to fetch all posts
func (h *Handler) GetPostsHandler(c *gin.Context) {
	posts, err := h.Queries.GetAllPosts(context.Background())
	if err != nil {
		h.Log.Errorf("Unable to fetch posts: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPostHandler handles GET requests to fetch a single post
func (h *Handler) GetPostHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Log.Errorf("Unable to convert id to int: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}
	post, err := h.Queries.GetPost(context.Background(), int32(id))
	if err != nil {
		h.Log.Errorf("Unable to fetch post: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch post"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// GetPostsByCategoryHandler handles GET requests to get all posts for a specific category
func (h *Handler) GetPostsByCategoryHandler(c *gin.Context) {
	postCategoryIDStr := c.Param("postCategoryID")
	postCategoryID, err := strconv.Atoi(postCategoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post category ID"})
		return
	}

	posts, err := h.Queries.GetPostsByCategory(context.Background(), pgtype.Int4{Int32: int32(postCategoryID), Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPostsByUserHandler handles GET requests to get all posts made by a specific user
func (h *Handler) GetPostsByUserHandler(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	posts, err := h.Queries.GetPostsByUser(context.Background(), pgtype.Int4{Int32: int32(userID), Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

type CreatePostApiParams struct {
	Title           string
	Content         string
	PostCategoryID  int
	AdditionalNotes string
}

// CreatePostHandler handles POST requests to create a new post
func (h *Handler) CreatePostHandler(c *gin.Context) {
	var req CreatePostApiParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface, userExists := c.Get("UserID")
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
		return
	}

	userID, userOk := userIDInterface.(int32)
	if !userOk {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
		return
	}

	err := h.Queries.CreatePost(context.Background(), db.CreatePostParams{
		Title:           req.Title,
		Content:         req.Content,
		UserID:          pgtype.Int4{Int32: userID, Valid: true},
		PostCategoryID:  pgtype.Int4{Int32: int32(1), Valid: true},
		AdditionalNotes: pgtype.Text{String: req.AdditionalNotes, Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

type UpdatePostApiParams struct {
	PostID          int
	Title           string
	Content         string
	PostCategoryID  int
	AdditionalNotes string
}

// UpdatePostHandler handles PUT requests to update an existing post
func (h *Handler) UpdatePostHandler(c *gin.Context) {
	var req UpdatePostApiParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface, userExists := c.Get("UserID")
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
		return
	}

	userID, userOk := userIDInterface.(int32)
	if !userOk {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
		return
	}

	err := h.Queries.UpdatePost(context.Background(), db.UpdatePostParams{
		PostID:          int32(req.PostID),
		Title:           req.Title,
		Content:         req.Content,
		UserID:          pgtype.Int4{Int32: userID, Valid: true},
		PostCategoryID:  pgtype.Int4{Int32: int32(req.PostCategoryID), Valid: true},
		AdditionalNotes: pgtype.Text{String: req.AdditionalNotes, Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

// DeletePostHandler handles DELETE requests to delete a post
func (h *Handler) DeletePostHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Log.Errorf("Unable to convert id to int: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	userIDInterface, userExists := c.Get("UserID")
	if !userExists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	userID, userOk := userIDInterface.(int32)
	if !userOk {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.Queries.DeletePostByPostIdAndUserId(context.Background(), db.DeletePostByPostIdAndUserIdParams{
		PostID: int32(id),
		UserID: pgtype.Int4{Int32: userID, Valid: true},
	})

	if err != nil {
		h.Log.Errorf("Unable to delete post: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

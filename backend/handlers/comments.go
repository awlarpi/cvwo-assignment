package handlers

import (
	"context"
	"net/http"
	"server/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateCommentApiParams struct {
	Content string
	PostID  int
	UserID  int
}

// CreateCommentHandler handles POST requests to create a new comment
func (h *Handler) CreateCommentHandler(c *gin.Context) {
	var req CreateCommentApiParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.validateUserID(c, int32(req.UserID)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	comment, err := h.Queries.CreateComment(context.Background(), db.CreateCommentParams{
		Content: req.Content,
		PostID:  pgtype.Int4{Int32: int32(req.PostID), Valid: true},
		UserID:  pgtype.Int4{Int32: int32(req.UserID), Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully", "comment": comment})
}

// GetCommentHandler handles GET requests to get a single comment by its ID
func (h *Handler) GetCommentHandler(c *gin.Context) {
	commentIDStr := c.Param("commentID")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	comment, err := h.Queries.GetComment(context.Background(), int32(commentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetCommentsByPostHandler handles GET requests to get all comments for a specific post
func (h *Handler) GetCommentsByPostHandler(c *gin.Context) {
	postIDStr := c.Param("postID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	comments, err := h.Queries.GetCommentsByPost(context.Background(), pgtype.Int4{Int32: int32(postID), Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetCommentsByUserHandler handles GET requests to get all comments made by a specific user
func (h *Handler) GetCommentsByUserHandler(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	comments, err := h.Queries.GetCommentsByUser(context.Background(), pgtype.Int4{Int32: int32(userID), Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// UpdateCommentHandler handles PUT requests to update a comment's content
func (h *Handler) UpdateCommentHandler(c *gin.Context) {
	var req db.UpdateCommentParams

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

	comment, err := h.Queries.UpdateCommentByCommentIdAndUserId(context.Background(), db.UpdateCommentByCommentIdAndUserIdParams{
		Content:   req.Content,
		CommentID: req.CommentID,
		UserID:    pgtype.Int4{Int32: userID, Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully", "comment": comment})
}

// DeleteCommentHandler handles DELETE requests to delete a comment by its ID
func (h *Handler) DeleteCommentHandler(c *gin.Context) {
	commentIDStr := c.Param("commentID")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
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

	err = h.Queries.DeleteCommentByCommentIdAndUserId(context.Background(), db.DeleteCommentByCommentIdAndUserIdParams{
		CommentID: int32(commentID),
		UserID:    pgtype.Int4{Int32: userID, Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

package controllers

import (
	"blog-platform/database"
	"blog-platform/models"
	"blog-platform/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not retrieve posts", nil)
		return
	}
	utils.CreateResponse(c, http.StatusOK, "Posts retrieved successfully", posts)
}

func GetPost(c *gin.Context) {
	postID := c.Param("post_id")

	var post models.Post
	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Post not found", nil)
		return
	}
	utils.CreateResponse(c, http.StatusOK, "Post retrieved successfully", post)
}

func CreatePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	post.AuthorID = user.(models.User).ID
	if err := database.DB.Create(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not create post", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Post created successfully", post)
}

func UpdatePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	currentUser := user.(models.User)

	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, "Invalid post ID", nil)
		return
	}

	var post models.Post
	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	if post.AuthorID != currentUser.ID {
		utils.CreateResponse(c, http.StatusForbidden, "You are not allowed to update this post", nil)
		return
	}

	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	if err := database.DB.Save(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not update post", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Post updated successfully", post)
}

func DeletePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	currentUser := user.(models.User)

	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, "Invalid post ID", nil)
		return
	}

	var post models.Post
	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	if post.AuthorID != currentUser.ID {
		utils.CreateResponse(c, http.StatusForbidden, "You are not allowed to delete this post", nil)
		return
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not delete post", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Post deleted successfully", nil)
}

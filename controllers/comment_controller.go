package controllers

import (
	"blog-platform/database"
	"blog-platform/models"
	"blog-platform/requests"
	"blog-platform/responses"
	"blog-platform/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCommentsByPost godoc
// @Summary Belirli bir posta ait yorumları getir
// @Description Post ID'ye göre yorumları getirir
// @Tags Comment
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} responses.CommentsResponse
// @Failure 404 {object} responses.ErrorResponse "Post bulunamadı"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /comments/post/{post_id} [get]
func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("post_id")

	var post models.Post
	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Post Not Found", nil)
		return
	}

	var comments []models.Comment
	if err := database.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not retrieve post comments", nil)
		return
	}

	var responseComments []responses.CommentResponse
	for _, comment := range comments {
		responseComments = append(responseComments, responses.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			AuthorID:  comment.AuthorID,
			PostID:    comment.PostID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	utils.CreateResponse(c, http.StatusOK, "Comments retrieved successfully", responses.CommentsResponse{Comments: responseComments})
}

// GetCommentsByUser godoc
// @Summary Kullanıcıya ait yorumları getir
// @Description Giriş yapmış kullanıcıya ait tüm yorumları getirir
// @Tags Comment
// @Produce json
// @Success 200 {object} responses.CommentsResponse
// @Failure 401 {object} responses.ErrorResponse "Yetkisiz"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /comments/user [get]
func GetCommentsByUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	currentUser := user.(models.User)
	var comments []models.Comment
	if err := database.DB.Where("author_id = ?", currentUser.ID).Find(&comments).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not retrieve user comments", nil)
		return
	}

	var responseComments []responses.CommentResponse
	for _, comment := range comments {
		responseComments = append(responseComments, responses.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			AuthorID:  comment.AuthorID,
			PostID:    comment.PostID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	utils.CreateResponse(c, http.StatusOK, "User comments retrieved successfully", responses.CommentsResponse{Comments: responseComments})
}

// CreateComment godoc
// @Summary Yeni bir yorum oluştur
// @Description Belirli bir posta yeni bir yorum ekler
// @Tags Comment
// @Accept json
// @Produce json
// @Param post_id path int true "Post ID"
// @Param comment body requests.CreateCommentRequest true "Yorum bilgisi"
// @Success 200 {object} responses.CommentResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 401 {object} responses.ErrorResponse "Yetkisiz"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /comments/{post_id} [post]
func CreateComment(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var input requests.CreateCommentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	postID, err := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if err != nil || postID == 0 {
		utils.CreateResponse(c, http.StatusBadRequest, "Valid Post ID is required", nil)
		return
	}

	comment := models.Comment{
		Content:  input.Content,
		AuthorID: user.(models.User).ID,
		PostID:   uint(postID),
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not create comment", nil)
		return
	}

	responseComment := responses.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		AuthorID:  comment.AuthorID,
		PostID:    comment.PostID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
	utils.CreateResponse(c, http.StatusOK, "Comment created successfully", responseComment)
}

// UpdateComment godoc
// @Summary Mevcut bir yorumu güncelle
// @Description Belirli bir yorumu günceller
// @Tags Comment
// @Accept json
// @Produce json
// @Param comment_id path int true "Comment ID"
// @Param comment body requests.UpdateCommentRequest true "Yorum bilgisi"
// @Success 200 {object} responses.CommentResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 401 {object} responses.ErrorResponse "Yetkisiz"
// @Failure 403 {object} responses.ErrorResponse "Erişim reddedildi"
// @Failure 404 {object} responses.ErrorResponse "Yorum bulunamadı"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /comments/{comment_id} [put]
func UpdateComment(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	commentID, err := strconv.ParseUint(c.Param("comment_id"), 10, 32)
	if err != nil || commentID == 0 {
		utils.CreateResponse(c, http.StatusBadRequest, "Valid Comment ID is required", nil)
		return
	}

	var input requests.UpdateCommentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var comment models.Comment
	if err := database.DB.Where("id = ?", uint(commentID)).First(&comment).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Comment not found", nil)
		return
	}
	if comment.AuthorID != user.(models.User).ID {
		utils.CreateResponse(c, http.StatusForbidden, "You are not allowed to update this comment", nil)
		return
	}

	comment.Content = input.Content
	if err := database.DB.Save(&comment).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not update comment", nil)
		return
	}

	responseComment := responses.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		AuthorID:  comment.AuthorID,
		PostID:    comment.PostID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
	utils.CreateResponse(c, http.StatusOK, "Comment updated successfully", responseComment)
}

// RemoveComment godoc
// @Summary Mevcut bir yorumu sil
// @Description Belirli bir yorumu siler
// @Tags Comment
// @Produce json
// @Param comment_id path int true "Comment ID"
// @Success 200 {object} responses.MessageResponse "Yorum başarıyla silindi"
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 401 {object} responses.ErrorResponse "Yetkisiz"
// @Failure 403 {object} responses.ErrorResponse "Erişim reddedildi"
// @Failure 404 {object} responses.ErrorResponse "Yorum bulunamadı"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /comments/{comment_id} [delete]
func RemoveComment(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	commentID, err := strconv.ParseUint(c.Param("comment_id"), 10, 32)
	if err != nil || commentID == 0 {
		utils.CreateResponse(c, http.StatusBadRequest, "Valid Comment ID is required", nil)
		return
	}

	var comment models.Comment
	if err := database.DB.Where("id = ?", uint(commentID)).First(&comment).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Comment not found", nil)
		return
	}
	if comment.AuthorID != user.(models.User).ID {
		utils.CreateResponse(c, http.StatusForbidden, "You are not allowed to delete this comment", nil)
		return
	}

	if err := database.DB.Delete(&comment).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not delete comment", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Comment deleted successfully", nil)
}

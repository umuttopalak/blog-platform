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

// GetPosts godoc
// @Summary Tüm postları getir
// @Description Tüm postları listeler
// @Tags Post
// @Produce json
// @Success 200 {object} responses.PostsResponse
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /posts [get]
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not retrieve posts", nil)
		return
	}

	var responsePosts []responses.PostResponse
	for _, post := range posts {
		responsePosts = append(responsePosts, responses.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			AuthorID:  post.AuthorID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	utils.CreateResponse(c, http.StatusOK, "Posts retrieved successfully", responses.PostsResponse{Posts: responsePosts})
}

// GetPost godoc
// @Summary Belirli bir postu getir
// @Description ID ile tek bir postu getirir
// @Tags Post
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} responses.PostResponse
// @Failure 404 {object} responses.ErrorResponse "Post bulunamadı"
// @Router /posts/{post_id} [get]
func GetPost(c *gin.Context) {
	postID := c.Param("post_id")

	var post models.Post
	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	responsePost := responses.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	utils.CreateResponse(c, http.StatusOK, "Post retrieved successfully", responsePost)
}

// CreatePost godoc
// @Summary Yeni bir post oluştur
// @Description Yeni bir post oluşturur
// @Tags Post
// @Accept json
// @Produce json
// @Param post body requests.CreatePostRequest true "Post bilgisi"
// @Success 200 {object} responses.PostResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var input requests.CreatePostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	post := models.Post{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: user.(models.User).ID,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not create post", nil)
		return
	}

	responsePost := responses.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	utils.CreateResponse(c, http.StatusOK, "Post created successfully", responsePost)
}

// UpdatePost godoc
// @Summary Mevcut bir postu güncelle
// @Description ID ile mevcut bir postu günceller
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path int true "Post ID"
// @Param post body requests.UpdatePostRequest true "Güncellenecek post bilgisi"
// @Success 200 {object} responses.PostResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 403 {object} responses.ErrorResponse "Yetkisiz erişim"
// @Failure 404 {object} responses.ErrorResponse "Post bulunamadı"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /posts/{post_id} [put]
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

	var input requests.UpdatePostRequest
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

	responsePost := responses.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	utils.CreateResponse(c, http.StatusOK, "Post updated successfully", responsePost)
}

// RemovePost godoc
// @Summary Mevcut bir postu sil
// @Description ID ile mevcut bir postu siler
// @Tags Post
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} responses.MessageResponse "Silme işlemi başarılı"
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 403 {object} responses.ErrorResponse "Yetkisiz erişim"
// @Failure 404 {object} responses.ErrorResponse "Post bulunamadı"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /posts/{post_id} [delete]
func RemovePost(c *gin.Context) {
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

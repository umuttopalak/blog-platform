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

// AddReaction godoc
// @Summary Yeni bir reaction ekle
// @Description Bir post veya yoruma reaction (like veya dislike) ekler
// @Tags Reaction
// @Accept json
// @Produce json
// @Param reaction body requests.CreateReactionRequest true "Reaction bilgisi"
// @Success 200 {object} responses.ReactionResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 401 {object} responses.ErrorResponse "Yetkisiz"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /reactions [post]
func AddReaction(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var input requests.CreateReactionRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Hem `PostID` hem de `CommentID` aynı anda sağlanırsa veya ikisi de sağlanmazsa hata döndür
	if (input.PostID == nil && input.CommentID == nil) || (input.PostID != nil && input.CommentID != nil) {
		utils.CreateResponse(c, http.StatusBadRequest, "You must provide either PostID or CommentID, not both", nil)
		return
	}

	reaction := models.Reaction{
		Type:      models.ReactionType(input.Type),
		UserID:    user.(models.User).ID,
		PostID:    input.PostID,
		CommentID: input.CommentID,
	}

	if err := database.DB.Create(&reaction).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not create reaction", nil)
		return
	}

	response := responses.ReactionResponse{
		ID:        reaction.ID,
		Type:      string(reaction.Type),
		UserID:    reaction.UserID,
		PostID:    reaction.PostID,
		CommentID: reaction.CommentID,
		CreatedAt: reaction.CreatedAt,
	}
	utils.CreateResponse(c, http.StatusOK, "Reaction added successfully", response)
}

// GetReactionsByPost godoc
// @Summary Belirli bir posta ait tüm reaction'ları getir
// @Description Belirli bir posta ait reaction'ları (like/dislike) listele
// @Tags Reaction
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} responses.ReactionsResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /reactions/post/{post_id} [get]
func GetReactionsByPost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, "Invalid Post ID", nil)
		return
	}

	var reactions []models.Reaction
	if err := database.DB.Where("post_id = ?", postID).Find(&reactions).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not retrieve reactions", nil)
		return
	}

	var responseReactions []responses.ReactionResponse
	for _, reaction := range reactions {
		responseReactions = append(responseReactions, responses.ReactionResponse{
			ID:        reaction.ID,
			Type:      string(reaction.Type),
			UserID:    reaction.UserID,
			PostID:    reaction.PostID,
			CommentID: reaction.CommentID,
			CreatedAt: reaction.CreatedAt,
		})
	}
	utils.CreateResponse(c, http.StatusOK, "Reactions retrieved successfully", responses.ReactionsResponse{Reactions: responseReactions})
}

// GetReactionsByComment godoc
// @Summary Belirli bir yoruma ait tüm reaction'ları getir
// @Description Belirli bir yoruma ait reaction'ları (like/dislike) listele
// @Tags Reaction
// @Produce json
// @Param comment_id path int true "Comment ID"
// @Success 200 {object} responses.ReactionsResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /reactions/comment/{comment_id} [get]
func GetReactionsByComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, "Invalid Comment ID", nil)
		return
	}

	var reactions []models.Reaction
	if err := database.DB.Where("comment_id = ?", commentID).Find(&reactions).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not retrieve reactions", nil)
		return
	}

	var responseReactions []responses.ReactionResponse
	for _, reaction := range reactions {
		responseReactions = append(responseReactions, responses.ReactionResponse{
			ID:        reaction.ID,
			Type:      string(reaction.Type),
			UserID:    reaction.UserID,
			PostID:    reaction.PostID,
			CommentID: reaction.CommentID,
			CreatedAt: reaction.CreatedAt,
		})
	}
	utils.CreateResponse(c, http.StatusOK, "Reactions retrieved successfully", responses.ReactionsResponse{Reactions: responseReactions})
}

// RemoveReaction godoc
// @Summary Bir reaction'ı sil
// @Description Belirli bir reaction'ı (like veya dislike) siler
// @Tags Reaction
// @Produce json
// @Param reaction_id path int true "Reaction ID"
// @Success 200 {object} responses.MessageResponse "Reaction silindi"
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 401 {object} responses.ErrorResponse "Yetkisiz"
// @Failure 403 {object} responses.ErrorResponse "Erişim reddedildi"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /reactions/{reaction_id} [delete]
func RemoveReaction(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	reactionID, err := strconv.Atoi(c.Param("reaction_id"))
	if err != nil || reactionID == 0 {
		utils.CreateResponse(c, http.StatusBadRequest, "Valid Reaction ID is required", nil)
		return
	}

	var reaction models.Reaction
	if err := database.DB.Where("id = ?", reactionID).First(&reaction).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Reaction not found", nil)
		return
	}

	if reaction.UserID != user.(models.User).ID {
		utils.CreateResponse(c, http.StatusForbidden, "You are not allowed to delete this reaction", nil)
		return
	}

	if err := database.DB.Delete(&reaction).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not delete reaction", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Reaction deleted successfully", nil)
}

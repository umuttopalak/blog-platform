package requests

// CreateCommentRequest yeni yorum oluşturmak için model
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// UpdateCommentRequest yorumu güncellemek için model
type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

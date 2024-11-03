package requests

type CreateReactionRequest struct {
	Type      string `json:"type" binding:"required,oneof=like dislike"`
	PostID    *uint  `json:"post_id,omitempty"`    // Sadece post için
	CommentID *uint  `json:"comment_id,omitempty"` // Sadece comment için
}

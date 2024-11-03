package responses

import (
	"time"

	"github.com/google/uuid"
)

// CommentResponse yorum bilgilerini temsil eden model
type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	AuthorID  uuid.UUID `json:"author_id"`
	PostID    uint      `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommentsResponse birden fazla yorumu temsil eden model
type CommentsResponse struct {
	Comments []CommentResponse `json:"comments"`
}

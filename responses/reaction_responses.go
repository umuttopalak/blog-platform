package responses

import (
	"time"

	"github.com/google/uuid"
)

type ReactionResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    *uint     `json:"post_id,omitempty"`
	CommentID *uint     `json:"comment_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ReactionsResponse struct {
	Reactions []ReactionResponse `json:"reactions"`
}

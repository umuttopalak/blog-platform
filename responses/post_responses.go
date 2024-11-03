package responses

import (
	"time"

	"github.com/google/uuid"
)

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostsResponse struct {
	Posts []PostResponse `json:"posts"`
}

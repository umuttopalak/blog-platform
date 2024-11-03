package models

import (
	"time"

	"github.com/google/uuid"
)

type ReactionType string

const (
	Like    ReactionType = "like"
	Dislike ReactionType = "dislike"
)

type Reaction struct {
	ID        uint         `gorm:"primaryKey"`
	UserID    uuid.UUID    `json:"user_id" gorm:"not null"`
	User      User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PostID    *uint        `json:"post_id,omitempty"`
	CommentID *uint        `json:"comment_id,omitempty"`
	Type      ReactionType `json:"type" gorm:"not null"`
	CreatedAt time.Time    `json:"created_at"`
}

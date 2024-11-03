package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CategoryID uint      `json:"category_id"`
	AuthorID   uuid.UUID `json:"author_id"`
	Author     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

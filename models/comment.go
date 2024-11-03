package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID    uint       `json:"post_id"`
	Post      Post       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	AuthorID  uuid.UUID  `json:"author_id"`
	Author    User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author,omitempty"`
	Reactions []Reaction `gorm:"foreignKey:CommentID" json:"reactions,omitempty"`
}

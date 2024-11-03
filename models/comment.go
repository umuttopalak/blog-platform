package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Like     int  `json:"like_count"`
	Dislike  int  `json:"dislike_count"`
	AuthorID uint `json:"author_id"`
	Author   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

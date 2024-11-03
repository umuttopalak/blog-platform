package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
}

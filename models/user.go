package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Posts     []Post    `gorm:"foreignKey:AuthorID"`
	Comments  []Comment `gorm:"foreignKey:AuthorID"`
}

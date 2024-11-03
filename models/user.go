package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string     `json:"first_name,omitempty"` // JSON çıktısında gösterilir
	LastName  string     `json:"last_name,omitempty"`  // JSON çıktısında gösterilir
	Username  string     `json:"username"`
	Email     string     `json:"-"` // JSON çıktısında gizlenir
	Password  string     `json:"-"` // JSON çıktısında gizlenir
	Posts     []Post     `gorm:"foreignKey:AuthorID" json:"-"`
	Comments  []Comment  `gorm:"foreignKey:AuthorID" json:"-"`
	Reactions []Reaction `gorm:"foreignKey:UserID" json:"-"`
}

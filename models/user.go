package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Posts     []Post     `gorm:"foreignKey:AuthorID" json:"-"`
	Comments  []Comment  `gorm:"foreignKey:AuthorID" json:"-"`
	Reactions []Reaction `gorm:"foreignKey:UserID" json:"-"`
	Roles     []Role     `gorm:"many2many:user_roles;" json:"-"`
}

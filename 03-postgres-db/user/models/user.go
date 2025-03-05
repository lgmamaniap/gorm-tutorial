package models

import (
	"github.com/lgmamaniap/gorm-tutorial/03-postgres-db/common/models"
	post "github.com/lgmamaniap/gorm-tutorial/03-postgres-db/user/post/models"
)

type User struct {
	models.Base
	Name  string      `json:"name" gorm:"size:255"`
	Email string      `json:"email" gorm:"size:255"`
	Posts []post.Post `json:"posts" gorm:"foreignKey:UserID"`
}

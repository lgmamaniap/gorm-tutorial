package models

import (
	"github.com/lgmamaniap/gorm-tutorial/03-postgres-db/common/models"
)

type Post struct {
	models.Base
	Title   string `json:"title" gorm:"size:255"`
	Content string `json:"content" gorm:"size:255"`
	UserID  uint64 `json:"user_id" gorm:"index"`
}

package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Article2 struct {
	ID uuid.UUID `json:"id"`
	Value  string    `json:"value"`
	UserID uuid.UUID `json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func GetArticle()  Article{
	var article Article
	if err := db.Find(&article).Error; err != nil {
		return article
	}else {
		return article
	}
}


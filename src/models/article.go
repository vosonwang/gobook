package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Article struct {
	Id        uuid.UUID `json:"id"`
	Article   string    `json:"article"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Article) TableName() string {
	return "nodes"
}

func FindArticle()  {
	
}
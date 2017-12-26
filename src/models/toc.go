package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Toc struct {
	ID uuid.UUID `json:"id"`
	Children string    `json:"children"`
	Lang     string    `json:"lang"`
	UserID   uuid.UUID `json:"author"`
	Name     string    `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}


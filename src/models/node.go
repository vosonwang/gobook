package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type node struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Expand    bool      `json:"expand"`
	Nodekey   int       `json:"nodeKey"`
	Active    bool      `json:"active"`
	Kind      int       `json:"kind"`
	ParentId  uuid.UUID `json:"parent_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type article struct {
	Id        uuid.UUID `json:"id"`
	Article   string    `json:"article"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Tocs []node

func GetNode(kind int) (Tocs,error){
	var tocs Tocs
	err:=db.Where("kind = ?", kind).Find(&tocs).Error
	return tocs,err
}
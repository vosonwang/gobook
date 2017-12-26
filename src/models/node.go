package models

import (
	"github.com/satori/go.uuid"
	"time"
	"io"
	"encoding/json"
)

type Node struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Expand    bool      `json:"expand"`
	NodeKey   int       `json:"nodeKey" gorm:"AUTO_INCREMENT"`
	Active    bool      `json:"active"`
	Kind      int       `json:"kind"`
	ParentId  uuid.UUID `json:"parent_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Article struct {
	Id        uuid.UUID `json:"id"`
	Article   string    `json:"article"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Tocs []Node

func GetNode(kind int) (Tocs, error) {
	var tocs Tocs
	err := db.Where("kind = ?", kind).Order("node_key").Find(&tocs).Error
	return tocs, err
}

func ParseNode(body io.Reader) (map[string]interface{}, error) {
	var a interface{}
	err := json.NewDecoder(body).Decode(&a)
	return a.(map[string]interface{}), err
}

func AddNode(node interface{}) (uuid.UUID, error) {
	a := node.(map[string]interface{})
	var b Node
	b.Id = uuid.NewV4()
	b.Title = a["title"].(string)
	b.ParentId, _ = uuid.FromString(a["parent_id"].(string))
	b.Kind = int(a["kind"].(float64))
	err := db.Create(&b).Error
	return b.Id, err
}

func FindNode(id string) (Node, error) {
	var node Node
	err := db.First(&node, "id=?", id).Error
	return node, err
}

func UpdateNode(node Node) error {
	err := db.Save(&node).Error
	return err
}

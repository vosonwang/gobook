package models

import (
	"github.com/satori/go.uuid"
	"time"
	"io"
	"encoding/json"
	"util"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func FindUser(user User) (User,bool) {

	a := db.Where(user).Find(&user)

	if a.RowsAffected == 1 {
		return user,true
	}

	return User{},false
}

func ParseUser(body io.Reader) User {
	var user User

	err := json.NewDecoder(body).Decode(&user)
	util.CheckErr(err)

	return user
}

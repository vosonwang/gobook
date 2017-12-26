package models

import (

	"fmt"
	"github.com/jinzhu/gorm"
	"util"
	"config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	/*Database*/

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.DbPort, config.User, config.Password, config.Dbname)
	var err error

	db, err = gorm.Open("postgres", psqlInfo)
	util.CheckErr(err)

	db.LogMode(true)

}


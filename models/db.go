package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/awesome?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(20)
	DB = db
}

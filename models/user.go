package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

func init() {
	if !DB.HasTable(&User{}) {
		if err := DB.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}

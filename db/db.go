package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	dsn := "root:NewdividE1@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		fmt.Println(err)
	}

	return db

}

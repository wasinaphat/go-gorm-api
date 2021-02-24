package models

import "github.com/go-gorm-api/db"

// type User struct {
// 	ID       int64  `json:"id"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Email    string `json:"email"`
// }
// type User struct {
// 	ID   int64  `json:"id"`
// 	Name string `json::"name"`
// 	Age  int64  `json:"age"`
// }

type User struct {
	ID   uint
	Name string
	Age  uint8
}

func init() {
	db.DB().Migrator().CreateTable(&User{})
}

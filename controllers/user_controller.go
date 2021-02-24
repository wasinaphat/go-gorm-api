package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gorm-api/db"
	"github.com/go-gorm-api/models"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	var User models.User
	id := c.Param("id")
	User.Name = "Tig"
	User.Age = 31
	user := User
	result := db.DB().Create(&user)
	fmt.Println(result.RowsAffected)

	c.JSON(http.StatusOK, gin.H{"id": id, "user": user})
}
func SignUp(c *gin.Context) {
	var User models.User
	id := c.Param("id")
	User.Name = "Tig"
	User.Age = 31
	user := User
	result := db.DB().Create(&user)
	fmt.Println(result.RowsAffected)

	c.JSON(http.StatusOK, gin.H{"id": id, "user": user})
}
func Login() {

}

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gorm-api/db"
	"github.com/go-gorm-api/models"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	var User models.User
	// id := c.Param("id")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result := db.DB().Where("id=?", id).Find(&User)
	fmt.Println("result.RowsAffected", result.RowsAffected)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Not found user id %d", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "user": User})
}
func SignUp(c *gin.Context) {
	var User models.User
	id := c.Param("id")
	// User.Name = "Tig"
	// User.Age = 31
	user := User
	result := db.DB().Create(&user)
	fmt.Println(result.RowsAffected)

	c.JSON(http.StatusOK, gin.H{"id": id, "user": user})
}
func Login() {

}

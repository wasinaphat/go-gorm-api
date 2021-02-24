package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	data, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal([]byte(data), &User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	result := db.DB().Where("username=?", User.Username).Find(&User)
	
	if result.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Already username %s", User.Username)})
		return
	} else {
		result := db.DB().Create(&User)
		if result.RowsAffected == 1 {
			c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("Sign up is successfully")})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Could not sign up")})
			return
		}
	}
	

}
func Login() {

}

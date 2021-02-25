package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	guard "github.com/go-gorm-api/guards"
)

var (
	router = gin.Default()
)

func AppRoute() {
	UserRoutes()
	v1 := router.Group("/v1")
	v1.Use(guard.AuthorizeJWT())
	{
		v1.GET("/test", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
	router.Run(":8080")
}

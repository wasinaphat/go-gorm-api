package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func AppRoute() {
	UserRoutes()
	router.Run(":8080")
}

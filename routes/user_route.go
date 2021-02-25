package routes

import "github.com/go-gorm-api/controllers"

func UserRoutes() {
	router.GET("/get-user/:id", controllers.GetUser)
	router.POST("/sign-up", controllers.SignUp)
	router.POST("/login", controllers.Login)
}

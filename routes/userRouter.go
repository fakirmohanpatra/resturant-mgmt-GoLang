package routes

import (
	"github.com/gin-gonic/gin"
	controllers "golang-resturant-mgmt/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.LogIn())
	incomingRoutes.PATCH("/users/:id", controllers.UpdateUser())
	incomingRoutes.DELETE("/users/:id", controllers.DeleteUser())
}
package routes

import (
	controllers "golang-resturant-mgmt/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/food", controllers.GetFood())
	incomingRoutes.GET("/food/:food_id", controllers.GetFood())
	incomingRoutes.POST("/food", controllers.CreateFood())
	incomingRoutes.PATCH("/food/:food_id", controllers.UpdateFood())
	incomingRoutes.DELETE("/food/:food_id", controllers.DeleteFood())
}

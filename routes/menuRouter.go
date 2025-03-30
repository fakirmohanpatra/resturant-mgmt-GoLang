package routes

import (
	controllers "golang-resturant-mgmt/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controllers.GetMenus())
	incomingRoutes.GET("/menu/:menu_id", controllers.GetMenu())
	incomingRoutes.POST("/menu", controllers.CreateMenu())
	incomingRoutes.PATCH("/menu/:menu_id", controllers.UpdateMenu())
	// incomingRoutes.DELETE("/menu/:menu_id", controllers.Delete())
}

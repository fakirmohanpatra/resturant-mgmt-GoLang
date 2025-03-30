package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"golang-resturant-mgmt/database"
	"golang-resturant-mgmt/models"
	"golang-resturant-mgmt/routes"
	"golang-resturant-mgmt/middlewares"
	"go.mongodb.org/mongo-driver/mongo"

)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port:= os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.useRoutes(router)
	router.Use(middlewares.Authentication())

	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.UserRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
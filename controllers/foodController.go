package controllers

import (
	"math"

	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get foods",
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get food",
		})
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create food",
		})
	}
}

func round(value float64) float64 {
	return float64(int(value*100)) / 100
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return round(num*output) / output
}

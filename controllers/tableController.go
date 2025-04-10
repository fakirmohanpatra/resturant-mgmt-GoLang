package controllers

import "github.com/gin-gonic/gin"

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all tables",
		})
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get table",
		})
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create table",
		})
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update table",
		})
	}
}

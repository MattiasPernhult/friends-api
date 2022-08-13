package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware function
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
	fmt.Println("hello world")
	fmt.Println("hello world 2")
}

func main() {
	// router := gin.Default()
	// router.Use(CORSMiddleware())
	// setUpRoutes(router)
	// router.Run("localhost:8080")
}

func setUpRoutes(router *gin.Engine) {
	// router.GET("/", controllers.IndexController)
	// router.GET("/persons", controllers.PersonsController)
	// router.GET("/persons/:name", controllers.PersonController)
	// router.GET("/episodes", controllers.EpisodesController)
	// router.GET("/overview", controllers.OverviewController)
}

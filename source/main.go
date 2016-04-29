package main

import (
	"sandbox/friends-api/source/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	setUpRoutes(router)
	router.Run("localhost:8080")
}

func setUpRoutes(router *gin.Engine) {
	router.StaticFile("/", "public/index.html")
	router.StaticFile("README.md", "README.md")
	router.GET("/persons", controllers.PersonsController)
}

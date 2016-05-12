package main

import (
	"friends-api/source/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	setUpRoutes(router)
	router.Run(":8080")
}

func setUpRoutes(router *gin.Engine) {
	router.GET("/", controllers.IndexController)
	router.GET("/persons", controllers.PersonsController)
	router.GET("/persons/:name", controllers.PersonController)
	router.GET("/episodes", controllers.EpisodesController)
}

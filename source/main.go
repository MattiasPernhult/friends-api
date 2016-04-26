package main

import (
	"sandbox/friends-api/source/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	setUpRoutes(router)
	router.Run()
}

func setUpRoutes(router *gin.Engine) {
	router.GET("people", controllers.GetPeople)
}

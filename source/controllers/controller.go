package controllers

import (
	"net/http"
	"sandbox/friends-api/source/services"

	"github.com/gin-gonic/gin"
)

// GetPeople function
func GetPeople(context *gin.Context) {
	people, err := services.GetPeople()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"people": people,
	})
}

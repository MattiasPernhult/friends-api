package controllers

import (
	"friends-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OverviewController function
func OverviewController(context *gin.Context) {
	overview, err := services.GetOverview()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"overview": overview,
		})
	}
}

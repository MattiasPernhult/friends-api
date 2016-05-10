package controllers

import (
	"net/http"
	customErrors "sandbox/friends-api/source/errors"
	"sandbox/friends-api/source/models"
	"sandbox/friends-api/source/services"

	"github.com/gin-gonic/gin"
)

func EpisodesController(context *gin.Context) {
	limitQuery := context.Query("limit")
	orderByQuery := context.Query("orderBy")

	eq := new(models.EpisodesQuery)

	if !eq.AddLimit(limitQuery) {
		context.JSON(customErrors.ErrorMessage(customErrors.ErrLimitParam))
		return
	}

	eq.AddOrderBy(orderByQuery)

	episodes, err := services.GetEpisodes(eq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"episodes": episodes,
	})
}

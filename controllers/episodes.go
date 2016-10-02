package controllers

import (
	customErrors "friends-api/errors"
	"friends-api/models"
	"friends-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EpisodesController function
func EpisodesController(context *gin.Context) {
	limitQuery := context.Query("limit")
	orderByQuery := context.Query("orderBy")

	eq := new(models.EpisodeQuery)

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

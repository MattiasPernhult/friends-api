package controllers

import (
	"net/http"
	customErrors "sandbox/friends-api/source/errors"
	"sandbox/friends-api/source/services"
	"sandbox/friends-api/source/utils"

	"github.com/gin-gonic/gin"
)

// PersonsController function
func PersonsController(context *gin.Context) {

	orderByQuery := context.Query("orderBy")
	limitQuery := context.Query("limit")

	limit, limitOK := utils.IsLimitParamValid(limitQuery)
	if !limitOK {
		context.JSON(customErrors.ErrorMessage(customErrors.ErrLimitParam))
		return
	}

	orderBy, orderByOK := utils.IsOrderByParamValid(orderByQuery)
	if !orderByOK {
		context.JSON(customErrors.ErrorMessage(customErrors.ErrPersonsOrderByParam))
		return
	}

	people, err := services.GetPersons(orderBy, limit)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"people": people,
	})
}

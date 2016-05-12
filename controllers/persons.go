package controllers

import (
	customErrors "friends-api/errors"
	"friends-api/models"
	"friends-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PersonsController function
func PersonsController(context *gin.Context) {

	limitQuery := context.Query("limit")
	orderByQuery := context.Query("orderBy")
	includeQuery := context.Query("include")

	pq := new(models.PersonsQuery)

	if !pq.AddLimit(limitQuery) {
		context.JSON(customErrors.ErrorMessage(customErrors.ErrLimitParam))
		return
	}

	pq.AddOrderBy(orderByQuery)
	pq.AddInclude(includeQuery)

	persons, err := services.GetPersons(pq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"persons": persons,
		})
	}
}

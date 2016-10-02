package controllers

import (
	"friends-api/models"
	"friends-api/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// PersonController function
func PersonController(context *gin.Context) {
	nameParam := context.Param("name")
	name := strings.ToLower(nameParam)

	pq := new(models.PersonQuery)
	pq.AddFind("accessName", name)

	persons, err := services.GetPerson(pq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"person": persons,
		})
	}
}

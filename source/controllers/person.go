package controllers

import (
	"net/http"
	"sandbox/friends-api/source/models"
	"sandbox/friends-api/source/services"
	"strings"

	"github.com/gin-gonic/gin"
)

// PersonController function
func PersonController(context *gin.Context) {
	nameParam := context.Param("name")
	name := strings.ToLower(nameParam)

	pq := new(models.PersonsQuery)
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
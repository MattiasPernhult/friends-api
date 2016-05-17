package services

import (
	"friends-api/errors"
	"friends-api/models"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// GetServices function
func GetOverview() (*models.Overview, *errors.RequestError) {
	var result models.Overview
	session, err := getSession()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &result, errors.ErrDatabaseConnection
	}

	err = session.
		DB("friends").
		C("overviews").
		Find(bson.M{}).
		One(&result)

	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &result, errors.ErrDatabaseQuery
	}

	return &result, nil
}

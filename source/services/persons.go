package services

import (
	"sandbox/friends-api/source/credentials"
	"sandbox/friends-api/source/errors"
	"sandbox/friends-api/source/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() (*mgo.Session, error) {
	mongoURI := credentials.GetCredential("mongoUri")
	session, err := mgo.Dial(mongoURI)
	return session, err
}

// GetPersons function
func GetPersons(pq *models.PersonsQuery) (*[]models.Persons, *errors.RequestError) {
	var results []models.Persons
	session, err := getSession()
	if err != nil {
		return &results, errors.ErrDatabaseConnection
	}

	err = session.
		DB("friends").
		C("people").
		Find(bson.M{}).
		Select(pq.Include).
		Sort(pq.OrderBy).
		Limit(pq.Limit).
		All(&results)

	if err != nil {
		return &results, errors.ErrDatabaseQuery
	}
	return &results, nil
}

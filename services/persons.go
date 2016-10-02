package services

import (
	"friends-api/errors"
	"friends-api/models"
	"log"
)

// GetPerson function
func GetPerson(pq *models.PersonQuery) (*models.Person, *errors.RequestError) {
	var result models.Person
	session, err := getSession()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &result, errors.ErrDatabaseConnection
	}

	err = session.
		DB("friends").
		C("people").
		Find(pq.Find).
		One(&result)

	if err != nil {
		log.Printf("ERROR: %v\n", err)
		log.Printf("Trying to find person: %v\n", pq.Find)
		return &result, errors.ErrDatabaseQuery
	}

	return &result, nil
}

// GetPersons function
func GetPersons(pq *models.PersonQuery) (*[]models.Person, *errors.RequestError) {
	var results []models.Person
	session, err := getSession()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &results, errors.ErrDatabaseConnection
	}

	err = session.
		DB("friends").
		C("people").
		Find(pq.Find).
		Select(pq.Include).
		Sort(pq.OrderBy...).
		Limit(pq.Limit).
		All(&results)

	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &results, errors.ErrDatabaseQuery
	}

	return &results, nil
}

package services

import (
	"log"
	"sandbox/friends-api/source/credentials"
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
func GetPersons(pq *models.PersonsQuery) (*[]models.Persons, error) {
	session, err := getSession()
	if err != nil {
		// TODO: Add error and return it
		log.Fatal(err)
	}

	var results []models.Persons

	log.Println("Inne i service")

	err = session.
		DB("friends").
		C("people").
		Find(bson.M{}).
		Select(pq.Include).
		Sort(pq.OrderBy).
		Limit(pq.Limit).
		All(&results)

	if err != nil {
		// TODO: Add error and return it
		log.Fatal(err)
	}
	return &results, err
}

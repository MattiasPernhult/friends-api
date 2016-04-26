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

// GetPeople function
func GetPeople() (*[]models.People, error) {
	session, err := getSession()
	if err != nil {
		log.Fatal(err)
	}

	var results []models.People

	err = session.
		DB("friends").
		C("people").
		Find(bson.M{}).
		Select(bson.M{"firstAppearance": 1, "lastAppearance": 1, "numberOfEpisodes": 1, "name": 1}).
		All(&results)

	if err != nil {
		log.Fatal(err)
	}
	return &results, err
}

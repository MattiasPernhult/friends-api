package services

import (
	"sandbox/friends-api/source/credentials"

	"gopkg.in/mgo.v2"
)

func getSession() (*mgo.Session, error) {
	mongoURI := credentials.GetCredential("mongoUri")
	session, err := mgo.Dial(mongoURI)
	return session, err
}

package services

import (
	"log"
	"sandbox/friends-api/source/errors"
	"sandbox/friends-api/source/models"
)

// GetEpisodes func
func GetEpisodes(eq *models.EpisodesQuery) (*[]models.Episodes, *errors.RequestError) {
	var results []models.Episodes

	session, err := getSession()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &results, errors.ErrDatabaseConnection
	}

	err = session.
		DB("friends").
		C("episodes").
		Find(eq.Find).
		Select(eq.Include).
		Sort(eq.OrderBy...).
		Limit(eq.Limit).
		All(&results)

	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &results, errors.ErrDatabaseQuery
	}

	return &results, nil
}

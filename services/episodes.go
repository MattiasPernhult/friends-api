package services

import (
	"friends-api/errors"
	"friends-api/models"
	"log"
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

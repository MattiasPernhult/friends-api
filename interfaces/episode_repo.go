package interfaces

import (
	"fmt"
	"friends-api/domain"
)

// EpisodeRepository struct
type EpisodeRepository struct {
	Logger    Logger
	DBHandler DBHandler
}

// Get function
func (repo *EpisodeRepository) Get(eq *domain.EpisodeQuery) ([]*domain.Episode, error) {

	orderByChar := []rune(eq.OrderBy[0])
	orderBy := toSnake(string(orderByChar[1:]))
	firstChar := orderByChar[0]
	if firstChar == '-' {
		orderBy += " DESC"
	}
	query := fmt.Sprintf("SELECT * FROM episodes LIMIT %d ORDER BY %s", eq.Limit, orderBy)
	repo.DBHandler.Query(query)
	return nil, nil
}

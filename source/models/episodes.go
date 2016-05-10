package models

import (
	"sandbox/friends-api/source/utils"

	"strings"
)

// Episodes struct
type Episodes struct {
	Season        string                `json:"season" bson:"season"`
	Episode       string                `json:"episode" bson:"episode"`
	SeasonNumber  int                   `json:"seasonNumber" bson:"seasonNumber"`
	EpisodeNumber int                   `json:"episodeNumber" bson:"episodeNumber"`
	Title         string                `json:"title" bson:"episodeTitle"`
	Airdate       string                `json:"airdate" bson:"airdate"`
	Conversations []EpisodeConversation `json:"conversations" bson:"conversations"`
	Plot          string                `json:"plot" bson:"plot"`
}

// EpisodeConversation struct
type EpisodeConversation struct {
	Said   string `json:"said" bson:"said"`
	Person string `json:"person" bson:"person"`
}

// EpisodesQuery struct
type EpisodesQuery struct {
	Limit   int
	OrderBy []string
	Find    map[string]interface{}
	Include map[string]interface{}
}

// AddLimit function
func (eq *EpisodesQuery) AddLimit(limitQuery string) bool {
	limit, ok := utils.IsLimitParamValid(limitQuery)
	eq.Limit = limit
	return ok
}

// AddOrderBy function
func (eq *EpisodesQuery) AddOrderBy(orderByQuery string) {
	orderBy := utils.IsOrderByParamValid(strings.Split(orderByQuery, ","),
		[]string{"seasonNumber", "-seasonNumber", "episodeNumber", "-episodeNumber"})
	eq.OrderBy = orderBy
}

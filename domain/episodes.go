package domain

import "strings"

// Episode struct
type Episode struct {
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

// EpisodeQuery struct
type EpisodeQuery struct {
	Limit   int
	OrderBy []string
	Find    map[string]interface{}
	Include map[string]interface{}
}

// AddFind function
func (eq *EpisodeQuery) AddFind(key string, value interface{}) {
	if eq.Find == nil {
		eq.Find = map[string]interface{}{}
	}
	eq.Find[key] = value
}

// AddLimit function
func (eq *EpisodeQuery) AddLimit(limitQuery string) bool {
	limit, ok := isLimitParamValid(limitQuery)
	eq.Limit = limit
	return ok
}

// AddOrderBy function
func (eq *EpisodeQuery) AddOrderBy(orderByQuery string) {
	orderBy := isOrderByParamValid(strings.Split(orderByQuery, ","),
		[]string{"seasonNumber", "-seasonNumber", "episodeNumber", "-episodeNumber"})
	eq.OrderBy = orderBy
}

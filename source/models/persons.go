package models

import (
	"sandbox/friends-api/source/utils"
	"strings"
)

// validIncludes
var validIncludes = []string{
	"firstAppearance",
	"lastAppearance",
	"numberOfEpisodes",
	"portrayedBy",
	"name",
	"nicknames",
	"gender",
	"dateOfBirth",
	"occupations",
	"relatives",
	"lines",
}

// Relative struct
type Relative struct {
	Name         string `json:"name,omitempty" bson:"name"`
	Relationship string `json:"relationship,omitempty" bson:"relationship"`
}

// CharacterLine struct
type CharacterLine struct {
	Person string `json:"name,omitempty" bson:"name"`
	Said   string `json:"said,omitempty" bson:"said"`
}

// Persons struct
type Persons struct {
	FirstAppearance  string     `json:"firstAppearance,omitempty" bson:"firstAppearance"`
	LastAppearance   string     `json:"lastAppearance,omitempty" bson:"lastAppearance"`
	NumberOfEpisodes int32      `json:"numberOfEpisodes,omitempty" bson:"numberOfEpisodes"`
	PotrayedBy       []string   `json:"portrayedBy,omitempty" bson:"portrayedBy"`
	Name             string     `json:"name,omitempty" bson:"name"`
	Nicknames        []string   `json:"nicknames,omitempty" bson:"nicknames"`
	Gender           string     `json:"gender,omitempty" bson:"gender"`
	DateOfBirth      string     `json:"dateOfBirth,omitempty" bson:"dateOfBirth"`
	Occupations      []string   `json:"occupations,omitempty" bson:"occupations"`
	Relatives        []Relative `json:"relatives,omitempty" bson:"relatives"`
	Lines            []string   `json:"lines,omitempty" bson:"lines"`
}

// PersonsQuery struct
type PersonsQuery struct {
	Find    map[string]interface{}
	Limit   int
	OrderBy string
	Include map[string]interface{}
}

// AddLimit function
func (pq *PersonsQuery) AddLimit(limitQuery string) bool {
	limit, ok := utils.IsLimitParamValid(limitQuery)
	pq.Limit = limit
	return ok
}

// AddOrderBy function
func (pq *PersonsQuery) AddOrderBy(orderByQuery string) bool {
	orderBy, ok := utils.IsOrderByParamValid(orderByQuery, []string{"-numberOfEpisodes", "numberOfEpisodes"})
	pq.OrderBy = orderBy
	return ok
}

// AddFind function
func (pq *PersonsQuery) AddFind(key string, value interface{}) {
	if pq.Find == nil {
		pq.Find = map[string]interface{}{}
	}
	pq.Find[key] = value
}

// AddInclude function
func (pq *PersonsQuery) AddInclude(includeQuery string) {
	include := map[string]interface{}{
		"firstAppearance":  1,
		"lastAppearance":   1,
		"numberOfEpisodes": 1,
		"name":             1,
	}
	pq.Include = utils.IsIncludeParamValid(strings.Split(includeQuery, ","), validIncludes, include)
}

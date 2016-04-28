package models

type (
	// Relatives struct
	Relatives struct {
		Name         string `json:"name,omitempty" bson:"name"`
		Relationship string `json:"relationship,omitempty" bson:"relationship"`
	}

	// Persons struct
	Persons struct {
		FirstAppearance  string      `json:"firstAppearance,omitempty" bson:"firstAppearance"`
		LastAppearance   string      `json:"lastAppearance,omitempty" bson:"lastAppearance"`
		NumberOfEpisodes int32       `json:"numberOfEpisodes,omitempty" bson:"numberOfEpisodes"`
		PotrayedBy       []string    `json:"portrayedBy,omitempty" bson:"portrayedBy"`
		Name             string      `json:"name,omitempty" bson:"name"`
		Nicknames        []string    `json:"nicknames,omitempty" bson:"nicknames"`
		Gender           string      `json:"gender,omitempty" bson:"gender"`
		DateOfBirth      string      `json:"dateOfBirth,omitempty" bson:"dateOfBirth"`
		Occupations      []string    `json:"occupations,omitempty" bson:"occupations"`
		Relatives        []Relatives `json:"relatives,omitempty" bson:"relatives"`
	}
)

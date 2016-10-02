package domain

// CreatedBy struct
type CreatedBy struct {
	Name  string `json:"name,omitempty" bson:"name"`
	Image string `json:"image,omitempty" bson:"image"`
}

// Season struct
type Season struct {
	StartAirdate     string `json:"startAirdate,omitempty" bson:"startAirdate"`
	NumberOfEpisodes int    `json:"numberOfEpisodes,omitempty" bson:"numberOfEpisodes"`
	PosterImage      string `json:"posterImage,omitempty" bson:"posterImage"`
	SeasonNumber     int    `json:"seasonNumber,omitempty" bson:"seasonNumber"`
}

// Overview struct
type Overview struct {
	HeaderImage      string      `json:"headerImage,omitempty" bson:"headerImage"`
	CreatedBy        []CreatedBy `json:"createdBy,omitempty" bson:"createdBy"`
	Language         string      `json:"language,omitempty" bson:"language"`
	Network          string      `json:"network,omitempty" bson:"network"`
	Country          string      `json:"country,omitempty" bson:"country"`
	Productions      []string    `json:"productions,omitempty" bson:"productions"`
	EpisodeRuntime   int         `json:"epsiodeRuntime,omitempty" bson:"epsiodeRuntime"`
	FirstAirdate     string      `json:"firstAirdate,omitempty" bson:"firstAirdate"`
	LastAirdate      string      `json:"lastAirdate,omitempty" bson:"lastAirdate"`
	NumberOfEpisodes int         `json:"numberOfEpisodes,omitempty" bson:"numberOfEpisodes"`
	NumberOfSeasons  int         `json:"numberOfSeasons,omitempty" bson:"numberOfSeasons"`
	Summary          string      `json:"summary,omitempty" bson:"overview"`
	PosterImage      string      `json:"posterImage,omitempty" bson:"posterImage"`
	Seasons          []Season    `json:"seasons,omitempty" bson:"seasons"`
}

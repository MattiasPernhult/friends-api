package domain

// EpisodeRepository interface
type EpisodeRepository interface {
	Get(eq *EpisodeQuery) ([]*Episode, error)
}

// PersonRepository interface
type PersonRepository interface {
	GetAll(pq *PersonQuery) ([]*Person, error)
	GetByName(pq *PersonQuery) (*Person, error)
}

// OverviewRepository interface
type OverviewRepository interface {
	Get() (*Overview, error)
}

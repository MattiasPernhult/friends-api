package interfaces

import "friends-api/domain"

// OverviewRepository function
type OverviewRepository struct {
	Logger Logger
}

// Get function
func (repo *OverviewRepository) Get() (*domain.Overview, error) {
	return nil, nil
}

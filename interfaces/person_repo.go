package interfaces

import "friends-api/domain"

// PersonRepository struct
type PersonRepository struct {
	Logger Logger
}

// GetAll function
func (repo *PersonRepository) GetAll(pq *domain.PersonQuery) ([]*domain.Person, error) {
	// TODO: add implementation
	return nil, nil
}

// GetByName function
func (repo *PersonRepository) GetByName(pq *domain.PersonQuery) (*domain.Person, error) {
	// TODO: add implementation
	return nil, nil
}

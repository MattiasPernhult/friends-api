package usecases

import "friends-api/domain"

// PersonInteractor struct
type PersonInteractor struct {
	PersonRepository domain.PersonRepository
	Logger           Logger
}

// GetAll function
func (interactor *PersonInteractor) GetAll(pq *domain.PersonQuery) ([]*domain.Person, error) {
	persons, err := interactor.PersonRepository.GetAll(pq)

	if err != nil {
		interactor.Logger.Log(err.Error())
	}

	return persons, err
}

// GetByName function
func (interactor *PersonInteractor) GetByName(pq *domain.PersonQuery) (*domain.Person, error) {
	person, err := interactor.PersonRepository.GetByName(pq)

	if err != nil {
		interactor.Logger.Log(err.Error())
	}

	return person, err
}

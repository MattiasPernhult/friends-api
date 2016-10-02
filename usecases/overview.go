package usecases

import "friends-api/domain"

// OverviewInteractor struct
type OverviewInteractor struct {
	OverviewRepository domain.OverviewRepository
	Logger             Logger
}

// Get function
func (interactor *OverviewInteractor) Get(*domain.Overview, error) {
	overview, err := interactor.OverviewRepository.Get()

	if err != nil {
		interactor.Logger.Log(err.Error())
	}

	return overview, err
}

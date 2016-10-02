package usecases

import "friends-api/domain"

// EpisodeInteractor struct
type EpisodeInteractor struct {
	EpisodeRepository domain.EpisodeRepository
	Logger            Logger
}

// Get function
func (interactor *EpisodeInteractor) Get(eq *domain.EpisodeQuery) ([]*Episode, error) {
	episodes, err := interactor.EpisodeRepository.Get(eq)

	if err != nil {
		interactor.Logger.Log(err.Error())
	}

	return episodes, err
}

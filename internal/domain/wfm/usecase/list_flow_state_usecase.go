package usecase

import (
	"cronos/internal/domain/wfm/entity"
	"cronos/internal/domain/wfm/repository"
)

type FlowStateUseCaseList struct {
	FlowStateListRepository *repository.FlowRepository
}

func (fr *FlowStateUseCaseList) Execute() ([]*entity.FlowState, error) {
	result, err := fr.FlowStateListRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

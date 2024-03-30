package usecase

import (
	"cronos/internal/domain/wfm/entity"
	"cronos/internal/domain/wfm/repository"
	"cronos/internal/dto"
)

type FlowSaveUseCase struct {
	FlowRepository repository.FlowRepository
}

func (fu *FlowSaveUseCase) Execute(flows *dto.FlowResponse) {
	if flows == nil || flows.Data == nil {
		return
	}

	var newFlowState []entity.FlowState
	for _, flow := range flows.Data {
		if flow.FlowState == nil {
			continue
		}
		for _, flowState := range flow.FlowState {
			newFlow := entity.NewFlow(entity.FlowState(flowState))
			if newFlow != nil {
				newFlowState = append(newFlowState, *newFlow)
			}
		}
	}
	for _, row := range newFlowState {
		if &row != nil {
			_ = fu.FlowRepository.UpdateOrInsert(row)
		}
	}
}

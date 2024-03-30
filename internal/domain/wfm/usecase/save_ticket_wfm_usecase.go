package usecase

import (
	"cronos/internal/domain/wfm/entity"
	repository2 "cronos/internal/domain/wfm/repository"
	"log"
)

type TicketSaveWFMUseCase struct {
	repository *repository2.WFMRepositoryImpl
}

func (t *TicketSaveWFMUseCase) Execute(tikects []entity.Ticket) error {
	for _, tikect := range tikects {
		newInput := entity.NewTicket(tikect)
		err := t.repository.UpsertTicket(newInput)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

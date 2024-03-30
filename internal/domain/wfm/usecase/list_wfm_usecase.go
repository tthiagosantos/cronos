package usecase

import (
	"cronos/configs"
	"cronos/internal/domain/wfm/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ListTicketWFMUseCase struct {
	FlowStatesUseCase FlowStatesUseCase
}

func (lt *ListTicketWFMUseCase) Execute(token string) (*[]entity.Ticket, *[]int, error) {
	url := configs.ConfigEnv.BASE_URL + "Ticket/GetMyTicketsApp?PageNumber=1&PageSize=1000"
	var idsTikects []int
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Erro ao criar a requisição HTTP:", err)
		return nil, nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação HTTP:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()
	var ticketResp entity.TicketResponse
	if err := json.NewDecoder(resp.Body).Decode(&ticketResp); err != nil {
		fmt.Println("Erro ao decodificar a resposta JSON:", err)
		return nil, nil, err
	}
	for _, tikect := range ticketResp.Data.Tickets {
		idsTikects = append(idsTikects, tikect.ID)
	}
	flowStatusResult, flowIds, err := lt.FlowStatesUseCase.Execute(idsTikects, token)
	if err != nil {
		return nil, nil, err
	}
	for i, ticket := range ticketResp.Data.Tickets {
		for _, flow := range flowStatusResult.Data {
			if ticket.ID == flow.TicketId {
				ticketResp.Data.Tickets[i].FlowStates = flow.FlowStates
				break
			}
		}
	}

	return &ticketResp.Data.Tickets, flowIds, nil
}

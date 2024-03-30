package usecase

import (
	"bytes"
	"cronos/configs"
	"cronos/internal/dto"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FlowStatesUseCase struct{}

func (lsu *FlowStatesUseCase) Execute(tikectsIds []int, token string) (*dto.FlowStateOutput, *[]int, error) {
	url := configs.ConfigEnv.BASE_URL + "Quiz/GetQuizByTicketsIdApp"
	jsonData, err := json.Marshal(tikectsIds)
	if err != nil {
		log.Println("Erro ao codificar JSON:", err)
		return nil, nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Erro ao criar a requisição HTTP:", err)
		return nil, nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição HTTP:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()
	var result dto.FlowStateOutput
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(err)
		return nil, nil, err
	}
	flowIds := []int{}
	for _, row := range result.Data {
		for _, rowFlow := range row.FlowStates {
			flowIds = append(flowIds, rowFlow.FlowStateId)
		}
	}
	return &result, &flowIds, nil
}

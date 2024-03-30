package usecase

import (
	"bytes"
	"cronos/configs"
	"cronos/internal/dto"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type FlowUseCase struct{}

func (fu *FlowUseCase) Execute(flowIds []int, token string) (*dto.FlowResponse, error) {
	uniqueIds := make(map[int]bool)
	var uniqueFlowIds []int
	for _, id := range flowIds {
		if !uniqueIds[id] {
			uniqueIds[id] = true
			uniqueFlowIds = append(uniqueFlowIds, id)
		}
	}
	url := configs.ConfigEnv.BASE_URL + "Flow/GetFlowlistApp"
	data := map[string]interface{}{
		"textToSearch": "string",
		"flowId":       uniqueFlowIds,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição HTTP:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response dto.FlowResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Println("Erro ao fazer o Unmarshal do JSON:", err)
		return nil, err
	}
	return &response, nil
}

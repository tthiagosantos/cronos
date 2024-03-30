package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type LoginResponse struct {
	Errors  []interface{} `json:"errors"`
	Error   interface{}   `json:"error"`
	Data    UserData      `json:"data"`
	IsValid bool          `json:"isValid"`
}

type UserData struct {
	ID                    int           `json:"id"`
	Name                  string        `json:"name"`
	PhotoURL              string        `json:"photoUrl"`
	Email                 string        `json:"email"`
	Token                 string        `json:"token"`
	TokenExpiresAt        string        `json:"tokenExpiresAt"`
	RefreshToken          string        `json:"refreshToken"`
	RefreshTokenExpiresAt string        `json:"refreshTokenExpiresAt"`
	UserLanguage          interface{}   `json:"userLanguage"`
	ProfileID             int           `json:"profileId"`
	ProfileName           string        `json:"profileName"`
	ProfileSettings       []interface{} `json:"profileSettings"`
}

type NewLoginUseCase struct{}

func (lg NewLoginUseCase) Execute() string {
	urlWfm := "https://i9wfm-web-dev.padtec.com.br/api/User/Login"
	jsonStr := strings.NewReader(`{"email": "admin@padtec.com.br","password": "$Padtec@2023!"}`)
	res, err := http.Post(urlWfm, "application/json", jsonStr)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return "Erro ao ler a resposta"
	}

	var result LoginResponse

	err = json.Unmarshal([]byte(responseBody), &result)
	if err != nil {
		fmt.Println("Erro ao fazer Unmarshal da resposta JSON:", err)
		return "Erro ao fazer Unmarshal da resposta JSON"
	}

	return result.Data.Token
}

package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CityRequest struct {
	City string `json:"city"`
}

// В ответе интересует лишь поле error. Если его значение false, значит город найден
type CityResponse struct {
	Error bool `json:"error"`
}

// Если переданное наименование города существует, вернёт объект CityRequest с полем city
func NewCityByString(inCity string) (*CityRequest, bool) {
	if blockBytes, err := json.Marshal(map[string]string{"city": inCity}); err == nil {
		resp, err := http.Post(
			"https://countriesnow.space/api/v0.1/countries/population/cities",
			"application/json",
			bytes.NewBuffer(blockBytes),
		)

		if err != nil {
			fmt.Println("Ошибка NewCityRequest/http.Post")
			return nil, false
		}

		if resp.StatusCode != 200 {
			fmt.Println("Ошибка NewCityRequest/StatusCode:", resp.StatusCode)
			return nil, false
		}

		respBodyBlock, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка NewCityRequest/io.ReadAll")
			return nil, false
		}

		var cityResponse CityResponse
		if errParse := json.Unmarshal(respBodyBlock, &cityResponse); errParse != nil {
			fmt.Println("Ошибка NewCityRequest/json.Unmarshal")
			return nil, false
		}

		if cityResponse.Error {
			fmt.Println("Ошибка NewCityRequest/ityResponse.Error == true")
			return nil, false
		}

		defer resp.Body.Close()

		return &CityRequest{City: inCity}, true
	}

	fmt.Println("Ошибка NewCityRequest/json.Marshal")
	return nil, false
}

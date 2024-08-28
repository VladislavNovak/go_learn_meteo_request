package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewCityRequestByIp() (*CityRequest, bool) {
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		fmt.Println("Ошибка NewCityRequestByIp/http.Get")
		return nil, false
	}

	if resp.StatusCode != 200 {
		fmt.Println("Ошибка NewCityRequestByIp/StatusCode:", resp.StatusCode)
		return nil, false
	}

	respBodyBlock, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка NewCityRequestByIp/io.ReadAll")
		return nil, false
	}

	var cityRequestByIp CityRequest
	if errParse := json.Unmarshal(respBodyBlock, &cityRequestByIp); errParse != nil {
		fmt.Println("Ошибка NewCityRequestByIp/json.Unmarshal")
		return nil, false
	}

	defer resp.Body.Close()

	return &cityRequestByIp, true
}

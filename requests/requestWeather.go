package requests

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func GetWeather(inCity *CityRequest, inFormat int) (string, bool) {
	if baseUrl, err := url.Parse("https://wttr.in/" + inCity.City); err == nil {
		// if baseUrl, err := url.Parse("https://wttr.in/"); err == nil {
		urlValues := url.Values{}
		// urlValues.Add("city", inCity.City)
		urlValues.Add("format", strconv.Itoa(inFormat))

		baseUrl.RawQuery = urlValues.Encode()

		resp, err := http.Get(baseUrl.String())
		if err != nil {
			fmt.Println("Ошибка GetWeather/http.Get")
			return "", false
		}

		if resp.StatusCode != 200 {
			fmt.Println("Ошибка GetWeather/StatusCode:", resp.StatusCode)
			return "", false
		}

		respBodyBlock, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка GetWeather/io.ReadAll")
			return "", false
		}

		resp.Body.Close()

		return string(respBodyBlock), true
	}

	return "", false
}

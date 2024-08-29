package main

import (
	"flag"
	"fmt"
	"learn/meteo_request/requests"
)

func main() {
	city := flag.String("city", "", "target town")
	format := flag.Int("format", 1, "weather info")
	flag.Parse()

	var cityRequest *requests.CityRequest
	var isCreate bool

	if *city != "" {
		cityRequest, isCreate = requests.NewCityByString(*city)
		if !isCreate {
			fmt.Println("Город не найден. Попробуйте снова")
			return
		}

	} else {
		cityRequest, isCreate = requests.NewCityRequestByIp()
		if !isCreate {
			fmt.Println("Местоположение не определено. Попробуйте снова")
			return
		}
	}

	weather, isWeather := requests.GetWeather(cityRequest, *format)

	if !isWeather {
		fmt.Printf("Невозможно получить погоду по указанному городу {%s}\n", cityRequest.City)
		return
	}

	fmt.Printf("Погода в городе {%s}: %s\n: ", cityRequest.City, weather)
}

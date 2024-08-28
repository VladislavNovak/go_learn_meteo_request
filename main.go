package main

import (
	"flag"
	"fmt"
	"learn/meteo_request/requests"
)

func main() {
	city := flag.String("city", "Moskva", "target town")
	flag.Parse()

	var cityRequest *requests.CityRequest
	var isCreate bool

	cityRequest, isCreate = requests.NewCityByString(*city)
	if !isCreate {
		fmt.Println("Прекращено 1")
		return
	}

	fmt.Println("Город по строке:", cityRequest.City)

	cityRequest, isCreate = requests.NewCityRequestByIp()
	if !isCreate {
		fmt.Println("Прекращено 2")
		return
	}

	fmt.Println("Город местный:", cityRequest.City)

	weather, isGet := requests.GetWeather(cityRequest, 1)
	if !isGet {
		fmt.Println("Прекращено 3")
		return
	}

	fmt.Println("Погода:", weather)
}

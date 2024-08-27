package main

import (
	"flag"
	"fmt"
	"learn/meteo_request/requests"
)

func main() {
	city := flag.String("city", "", "target town")
	flag.Parse()
	result, isCreate := requests.NewCityRequest(*city)
	if !isCreate {
		fmt.Println("Прекращено")
		return
	}
	fmt.Println("Город:", result.City)
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "target town")
	flag.Parse()
	fmt.Println("Вы ввели: ", *city)
}

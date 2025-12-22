package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	Parking := map[string]int{
		"A1": 300,
		"A2": 350,
		"A3": 340,
		"A4": 370,
		"A5": 400,
		"B1": 910,
		"B2": 980,
		"B3": 900,
		"B4": 1000,
		"B5": 1100,
	}

	for k, _ := range Parking {
		if Parking[k] < 500 {
			fmt.Println("Место:", k)
		}
	}

	pp.Println("Парковка до учета скидок:", Parking)
	for k, _ := range Parking {
		if Parking[k] > 900 {
			Parking[k] = int(float64(Parking[k]) * 0.9)
		}
	}
	pp.Println("Парковка после учета скидок:", Parking)

}

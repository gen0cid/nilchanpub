package main

import "fmt"

type Car struct {
	Name          string
	EngineVolume  float64
	HP            int
	NumberOfSeats int
	isPetrol      bool
}

func main() {
	car1 := Car{
		Name:          "Shevrolet",
		EngineVolume:  5.5,
		HP:            515,
		NumberOfSeats: 6,
		isPetrol:      true,
	}

	car2 := Car{
		Name:          "Mersedes-Benz",
		EngineVolume:  4.0,
		HP:            510,
		NumberOfSeats: 4,
		isPetrol:      true,
	}

	car3 := Car{
		Name:          "BMW",
		EngineVolume:  4.4,
		HP:            727,
		NumberOfSeats: 4,
		isPetrol:      true,
	}

	fmt.Println("Car1:", car1)
	fmt.Println("Car2:", car2)
	fmt.Println("Car3:", car3)
	fmt.Println()

	car1.Name = "Toyota"
	car1.EngineVolume = 4.5
	car1.HP = 535

	fmt.Println("Car1:", car1)
	fmt.Println("Car2:", car2)
	fmt.Println("Car3:", car3)
}

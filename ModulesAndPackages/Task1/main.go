package main

import (
	car "Task1/Car"

	"github.com/k0kubun/pp"
)

func main() {
	car1 := car.NewCar("Cadillac", "Escalade", 416, 6.2)

	pp.Println("car1:", car1)

}

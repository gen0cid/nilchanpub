package figures

import "fmt"

type Circle struct {
	diameter float64
}

func NewCircle(diameter float64) Circle {
	return Circle{
		diameter: diameter,
	}
}

func (c Circle) CountingSquare() float64 {
	fmt.Println("я окружносить и вот моя площадь")
	return c.diameter * 3.14

}

package figures

import "fmt"

type Square struct {
	length float64
	width  float64
}

func NewSquare(length float64, width float64) Square {
	return Square{
		length: length,
		width:  width,
	}
}

func (s Square) CountingSquare() float64 {
	fmt.Println("Я квадрат, со сторонами ", s.length, " и ", s.width, ". Вот моя площадь: ", s.length*s.width)
	return s.length * s.width

}

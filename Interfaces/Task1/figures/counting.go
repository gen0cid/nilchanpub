package figures

import "fmt"

type figures interface {
	CountingSquare() float64
}

func CountingSquare(f figures) {
	fmt.Println("Считаю площадь фигуры: ")
	fmt.Println(f.CountingSquare())

}

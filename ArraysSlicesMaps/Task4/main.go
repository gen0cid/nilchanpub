package main

import "fmt"

type Dogs struct {
	Name     string
	Rating   int
	IsPolite bool
}

func main() {
	dogs := []Dogs{
		{
			Name:     "Матильда",
			Rating:   6,
			IsPolite: true,
		},
		{
			Name:     "Шарик",
			Rating:   9,
			IsPolite: true,
		},
		{
			Name:     "Бобик",
			Rating:   5,
			IsPolite: false,
		},
	}
	fmt.Println("Собаки до:", dogs)
	for i, _ := range dogs {
		if dogs[i].IsPolite {
			dogs[i].Rating++
		}
	}
	fmt.Println("Собаки после:", dogs)

}

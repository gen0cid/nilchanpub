package main

import "fmt"

func main() {
	beerDelicious := true
	croutonsDelicious := true

	if beerDelicious && croutonsDelicious {
		fmt.Println("Мы пойдем гулять")
	} else {
		fmt.Println("Мы никуда НЕ идем")
	}
}

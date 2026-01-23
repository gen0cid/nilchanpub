package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func postMan(transferPoint chan string) {

	letters := []rune("abcdefghijklmnopqrstuvwxyz")

	iterations := 1 + rand.IntN(10)
	for i := 1; i <= iterations; i++ {
		interval := 300 + rand.IntN(400)
		time.Sleep(time.Duration(interval) * time.Millisecond)

		slice := make([]rune, 5)

		for i := range slice {
			slice[i] = letters[rand.IntN(len(letters))]
		}

		transferPoint <- string(slice)
	}
	close(transferPoint)
}
func main() {
	transferPoint := make(chan string)
	feedbacks := []string{}

	go postMan(transferPoint)

	for i := range transferPoint {
		feedbacks = append(feedbacks, i)
	}

	fmt.Println("отзывы:", feedbacks)
}

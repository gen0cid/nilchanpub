package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

var reviews []string
var mtx sync.Mutex

func review(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		letters := []rune("abcdefghijklmnopqrstuvwxyz")
		onerew := []rune{}
		for i := 0; i < 3; i++ {
			onerew = append(onerew, letters[rand.IntN(len(letters))])
		}
		mtx.Lock()
		reviews = append(reviews, string(onerew))
		mtx.Unlock()
	}
}

func main() {
	quantity := 1 + rand.IntN(100)
	wg := &sync.WaitGroup{}
	fmt.Println("Количество фанатов:", quantity)

	for i := 0; i < quantity; i++ {
		wg.Add(1)
		go review(wg)
	}

	wg.Wait()

	mtx.Lock()
	fmt.Println(reviews)
	mtx.Unlock()

	mtx.Lock()
	fmt.Println("длинна rewews", len(reviews))
	mtx.Unlock()
	fmt.Println("завершение...")
}

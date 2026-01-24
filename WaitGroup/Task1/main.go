package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func gadener(wg *sync.WaitGroup, n int) {
	defer wg.Done()

	beds := 1 + rand.IntN(5)
	fmt.Printf("Я огородник №%d, мне надо полить: %d грядок!\n", n, beds)

	for i := 1; i <= beds; i++ {
		fmt.Printf("Огородник №%d поливает грядку под номером %d\n", n, i)
		time.Sleep(time.Duration(500+rand.IntN(501)) * time.Millisecond)
	}

}

func main() {
	gardenersCount := 3 + rand.IntN(5)

	fmt.Printf("Запускаем %d огородников\n", gardenersCount)

	wg := &sync.WaitGroup{}

	wg.Add(gardenersCount)
	for i := 1; i <= gardenersCount; i++ {
		go gadener(wg, i)

	}

	wg.Wait()
	fmt.Println("Я main и я завершился")
}

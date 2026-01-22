package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			fmt.Println("Привет, я иду в зал!!!")
			time.Sleep(100 * time.Millisecond)
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			fmt.Println("Привет, я учу golang!!!")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			fmt.Println("Привет, какой хороший день!!!")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()

}

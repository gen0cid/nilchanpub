package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var number int
var Number atomic.Int64
var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mtx.Lock()
		number++
		mtx.Unlock()
	}
}

func Increase(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		Number.Add(1)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(20)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go Increase(wg)
	go Increase(wg)
	go Increase(wg)
	go Increase(wg)
	go Increase(wg)
	go Increase(wg)
	go Increase(wg)
	go Increase(wg)
	go Increase(wg)
	go Increase(wg)

	wg.Wait()
	fmt.Println("number:", number)

	fmt.Println("Number:", Number.Load())
}

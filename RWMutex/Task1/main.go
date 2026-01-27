package main

import (
	"fmt"
	"sync"
	"time"
)

var users = map[string]string{}
var mtx sync.RWMutex

func write(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000000; i++ {
		mtx.Lock()
		users[fmt.Sprintf("%d", i)] = fmt.Sprintf("%d", i)
		mtx.Unlock()
	}

}
func get(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		mtx.RLock()
		_ = users[fmt.Sprintf("%d", i)]
		mtx.RUnlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	timeNow := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go get(wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write(wg)
	}
	wg.Wait()

	fmt.Println(time.Since(timeNow))
}

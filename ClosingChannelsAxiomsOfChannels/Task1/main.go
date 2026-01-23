package main

import "fmt"

// closed channel

func closedChannelClose() {
	// panic: close of closed channel

	ch := make(chan int)

	close(ch)
	close(ch)
}

func closedChannelRead() {
	// v: 0
	// ok: false
	// можем читать из закрытого канала -> получим значение по умолчанию
	// переменная ok показывает нам полученное значение(v) по умолчанию(false) или нет(true)

	ch := make(chan int)
	close(ch)

	v, ok := <-ch
	fmt.Println("v:", v)
	fmt.Println("ok:", ok)
}

func closedChannelWrite() {
	// panic: send on closed channel

	ch := make(chan int)
	close(ch)

	ch <- 10
}

// nil channel

func nilChannelClose() {
	// panic: close of nil channel

	var ch chan string
	close(ch)
}

func nilChannelRead() {
	// fatal error: all goroutines are asleep - deadlock!

	var ch chan string
	go func() {
		ch <- "hello"
	}()
	v := <-ch
	fmt.Println("v:", v)
}

func nilChannelWrite() {
	// fatal error: all goroutines are asleep - deadlock!

	var ch chan string

	ch <- "hello"
}

func main() {
	nilChannelRead()
}

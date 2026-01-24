package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("foo завершилсь :(", n)
			return
		default:
			fmt.Println("Привет, я foo!!!", n)
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func boo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("boo завершилсь :(", n)
			return
		default:
			fmt.Println("Привет, я boo!!!", n)
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func loo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("loo завершилсь :(", n)
			return
		default:
			fmt.Println("Привет, я loo!!!", n)
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func main() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	middleCtx, middleCancel := context.WithCancel(parentCtx)
	childCtx, childCancel := context.WithCancel(middleCtx)

	for i := 1; i <= 3; i++ {
		go foo(parentCtx, i)
	}

	for i := 1; i <= 3; i++ {
		go boo(middleCtx, i)
	}

	for i := 1; i <= 3; i++ {
		go loo(childCtx, i)
	}

	time.Sleep(3 * time.Second)

	childCancel()

	time.Sleep(1 * time.Second)

	middleCancel()

	time.Sleep(1 * time.Second)

	parentCancel()

	time.Sleep(3 * time.Second)

	fmt.Println("я main и я завершился)))")
}

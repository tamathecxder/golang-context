package golangcontext

import (
	"fmt"
	"runtime"
	"testing"
)

func CreateCounterLeak() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0

		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func TestGoroutineLeak(t *testing.T) {
	fmt.Println("before, total goroutine:", runtime.NumGoroutine())

	destination := CreateCounterLeak()

	for n := range destination {
		fmt.Println("Counter:", n)

		if n == 10 {
			break
		}
	}

	fmt.Println("after, total goroutine:", runtime.NumGoroutine())
}

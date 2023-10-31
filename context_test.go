package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	fmt.Println("before, total goroutine:", runtime.NumGoroutine())

	destination := CreateCounter(ctx)

	fmt.Println("process, total goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)

		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(3 * time.Second)

	fmt.Println("after, total goroutine:", runtime.NumGoroutine())
}

func TestContextValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println("-----------------------------------")

	fmt.Println(contextF.Value("f")) // success
	fmt.Println(contextF.Value("c")) // success, because C is F parent
	fmt.Println(contextF.Value("b")) // fail, the parent is different
	fmt.Println(contextA.Value("e")) // fail, parent cannot access value from its child
}

func TestContext(t *testing.T) {
	background := context.Background()

	fmt.Println(background)

	todo := context.TODO()

	fmt.Println(todo)
}

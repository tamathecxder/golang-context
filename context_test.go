package golangcontext

import (
	"context"
	"fmt"
	"testing"
)

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

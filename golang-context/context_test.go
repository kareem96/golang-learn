package golangcontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T)  {
	background := context.Background()
	fmt.Println("this bacground : ", background)

	todo := context.TODO()
	fmt.Println("this bacground : ", todo)
}


func TestContextWithValue(t *testing.T)  {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")


	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")


	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b")) // nil

	fmt.Println(contextA.Value("b")) // nil
}

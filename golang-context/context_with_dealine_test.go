package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterDeadline(ctx context.Context) chan int  {

	// func CreateCounter() chan int  {
	
		destination := make(chan int)
		go func() {
			defer close(destination)
			counter := 1
	
			for {
				select {
				case <- ctx.Done(): return
				default:
					destination <- counter
					counter++
					time.Sleep(1 * time.Second)
				}
			}
	
			// without context
			// for {
			// 	destination <- counter
			// 	counter++
			// }
	
		}()
	
		return destination
	}
	
	func TestContextWithDeadline(t *testing.T)  {
	
		fmt.Println("Total Goroutine", runtime.NumGoroutine())
	
		// adding context with cancel
		parent := context.Background()
		ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))

		defer cancel()
	
		
		// destination := CreateCounter()
	
		destination := CreateCounterDeadline(ctx)
	
		fmt.Println("Total Goroutine", runtime.NumGoroutine())
	
		for n := range destination {
			fmt.Println("Counter", n)
			// if n == 10 {
			// 	break
			// }
		}
		
		time.Sleep(2 * time.Second)
	
		fmt.Println("Total Goroutine", runtime.NumGoroutine())
	}
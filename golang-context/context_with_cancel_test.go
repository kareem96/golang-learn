package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// with context cancel
func CreateCounter(ctx context.Context) chan int  {

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
	
	func TestContextWithCancel(t *testing.T)  {
	
		fmt.Println("Total Goroutine", runtime.NumGoroutine())
	
		// adding context with cancel
		parent := context.Background()
		ctx, cancel := context.WithCancel(parent)
	
		
		// destination := CreateCounter()
	
		destination := CreateCounter(ctx)
	
		fmt.Println("Total Goroutine", runtime.NumGoroutine())
	
		for n := range destination {
			fmt.Println("Counter", n)
			if n == 10 {
				break
			}
		}
	
		cancel() // mengirim signal cancel ke context
		time.Sleep(2 * time.Second)
	
		fmt.Println("Total Goroutine", runtime.NumGoroutine())
	}
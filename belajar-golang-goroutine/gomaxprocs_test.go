package belajargolanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)


func TestGomaxProcs(t *testing.T)  {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU : ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)

	group.Wait()
}


func TestChangeNumberThread(t *testing.T)  {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU : ", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)

	group.Wait()
}
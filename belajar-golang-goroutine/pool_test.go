package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T)  {
	pool := sync.Pool{
		//default if jika return nil change to new
		New: func () interface{}  {
			return "New"
		},
	}

	pool.Put("Abdul")
	pool.Put("Karim")
	pool.Put("Melayu")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Selesai")
}
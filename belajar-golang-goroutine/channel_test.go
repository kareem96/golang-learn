package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreatingChannel(t *testing.T)  {
	channel := make(chan string)

	go func(){
		time.Sleep(2 * time.Second)
		channel <- "Abdul Karim"
		fmt.Println("Selesai mengritim channel")
	}()
	data := <- channel
	fmt.Println(data)

	defer close(channel)
}


//channel with parameter
func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	close(channel)
}

func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Abdul Karim"
}


// channel in
func OnlyIn(channel chan <- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Abdul Karim"
}

// channel out
func OnlyOut(channel <- chan string)  {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T)  {
	channel := make(chan string)
	
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}


// buffered channel
func TestBufferedChannel(t * testing.T)  {
	// adding buffered number
	channel := make(chan string, 2)

	defer close(channel)

	// channel <- "Kareem"
	// fmt.Println(<- channel)

	go func() {
		channel <- "Abdul"
		channel <- "Karim"
	}()

	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Selesai")
}


// range channel
func TestRangeChannel(t *testing.T)  {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++{
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data " + data)
	}
}


// select channel
func TestSelectChannel(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// select{
	// case data := <- channel1:
	// 	fmt.Println("Data dari channel 1", data)
	// case data := <- channel2:
	// 	fmt.Println("Data dari channel 2", data)
	// }

	counter := 0
	for {
		select{
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2{
			break
		}
	}	
}


// default select channel 
func TestSelectChannelDefault(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select{
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			counter++

		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
//
		default: 
		fmt.Println("Menunggu Data")

		}
		if counter == 2{
			break
		}
	}	
}



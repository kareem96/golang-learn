package main

import "fmt"

func main()  {
	name := "kareem"

	if name == "kareem"{
		fmt.Println("Hello")
	} else if name == "abdul"{
		fmt.Println("Hello abdul")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	length := len(name)
	// if length > 5 {
	// 	fmt.Println("nama terlalu panjang")
	// }else {
	// 	fmt.Println("nama sudah benar")
	// }

	//short statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	}else {
		fmt.Println("nama sudah benar")
	}

	//siwthc without condition
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah bener")
	}
	
}
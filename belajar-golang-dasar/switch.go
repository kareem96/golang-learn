package main

import "fmt"

func main() {
	name := "kareem"

	switch name {
	case "kareem":
		fmt.Println("Hello kareem")
	case "abdul":
		fmt.Println("Hello abdul")
	}
	length := len(name);
	switch length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
		
	}

	//switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
		
	}


}
package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("Youre blocked", name)
	}else{
		fmt.Println("welcome", name)
	}
}

func main() {
	blacklist := func (name string) bool {
		return name == "anjing"
	}

	registerUser("kareem", blacklist)

	//another call
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
package main

import "fmt"

// function type declarations
type Filter func(string) string


// before
// func sayHelloWithFilter(name string, filter func(string) string) {

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("kareem", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}

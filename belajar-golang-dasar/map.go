package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "kareem"
	// person["address"] = "jakarta"

	// fmt.Println(person)

	person := map[string]string{
		"name":    "kareem",
		"address": "jakarta",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)


	book := make(map[string]string)
	book["title"] = "Book of Golang"
	book["author"] = "kareem"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}

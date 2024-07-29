package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct{
	FirstName,MiddleName,LastName, Gender string
	Age int
	Hobbies []string
	Addresses []Address
}

type Address struct{
	Street string
	Country string
	PostalCode string
}

func TestJsonObject(t *testing.T)  {
	customer := Customer{
		FirstName: "Abdul",
		MiddleName: "Karim",
		LastName: "Melayu",
		Gender: "Pria",
		Age: 28,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

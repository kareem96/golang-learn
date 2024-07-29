package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)


func TestStreamEncoder(t *testing.T)  {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Abdul",
		MiddleName: "Karim",
		LastName: "Melayu",
	}
	_ = encoder.Encode(customer)

	fmt.Println(customer)
}
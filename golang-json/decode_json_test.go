package main

import (
	"encoding/json"
	"fmt"
	"testing"
)


func TestDecodeJson(t *testing.T)  {
	jsonRequest := `{
		"FirstName":"Abdul",
		"MiddleName":"Karim",
		"LastName":"Melayu",
		"Gender":"Male",
		"Age":28
	}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil{
		panic(err)
	}
	
	fmt.Println(customer)
}
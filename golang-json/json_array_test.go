package main

import (
	"encoding/json"
	"fmt"
	"testing"
)


func TestJsonArray(t *testing.T)  {
	customer := Customer{
		FirstName: "Abdul",
		MiddleName: "Karim",
		LastName: "Melayu",
		Gender: "Pria",
		Age: 28,
		Hobbies: []string{ // this json slice/array string
			"Gaming",
			"Coding",
			"Reading",
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T)  {
	jsonString := `{"FirstName":"Abdul","MiddleName":"Karim","LastName":"Melayu","Gender":"Pria","Age":28,"Hobbies":["Gaming","Code","Read"]}`
	jsonByte := []byte(jsonString)
	
	customer := &Customer{}

	err := json.Unmarshal(jsonByte,customer)
	if err != nil{
		panic(err)
	}
	fmt.Println(customer)
}


func TestJsonArrayComplex(t *testing.T)  {
	customer := Customer{
		FirstName: "Abdul",
		MiddleName: "Karim",
		LastName: "Melayu",
		Gender: "Pria",
		Age: 28,
		Hobbies: []string{ // this json slice/array string
			"Gaming",
			"Coding",
			"Reading",
		},
		Addresses: []Address{ // this json slice/array complex
			{
				Street: "Jalan1",
				Country: "Country1",
				PostalCode: "Postal1",
			},
			{
				Street: "Jalan2",
				Country: "Country2",
				PostalCode: "Postal2",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T)  {
	jsonString := `{"FirstName":"Abdul","MiddleName":"Karim","LastName":"Melayu","Gender":"Pria","Age":28,"Hobbies":["Gaming","Coding","Reading"],"Addresses":[{"Street":"Jalan1","Country":"Country1","PostalCode":"Postal1"},{"Street":"Jalan2","Country":"Country2","PostalCode":"Postal2"}]}`
	jsonByte := []byte(jsonString)


	customer := &Customer{}

	err := json.Unmarshal(jsonByte,customer)
	if err != nil{
		panic(err)
	}
	fmt.Println(customer)
}
func TestOnlyJsonArrayComplexDecode(t *testing.T)  {
	jsonString := `[{"Street":"Jalan1","Country":"Country1","PostalCode":"Postal1"},{"Street":"Jalan2","Country":"Country2","PostalCode":"Postal2"}]`
	jsonByte := []byte(jsonString)


	addresses := &[]Address{}

	err := json.Unmarshal(jsonByte,addresses)
	if err != nil{
		panic(err)
	}
	fmt.Println(addresses)
}

func TestOnlyJsonArrayComplex(t *testing.T)  {
	
		addresses := []Address{ // this json slice/array complex
			{
				Street: "Jalan1",
				Country: "Country1",
				PostalCode: "Postal1",
			},
			{
				Street: "Jalan2",
				Country: "Country2",
				PostalCode: "Postal2",
			},
		}
	

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
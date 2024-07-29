package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{})  {
	bytes, err := json.Marshal(data)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMarsahl(t *testing.T)  {
	logJson("Kareem")
	logJson(1)
	logJson("true")
	logJson([]string{
		"abdul",
		"karim",
		"melayu",
	})
}
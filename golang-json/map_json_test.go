package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJson(t *testing.T)  {
	jsonString := `{"id":"001","name":"Macbook","price":"10USD","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	fmt.Println(result["image_url"])
}
func TestMapJsonEncode(t *testing.T)  {
	product := map[string] interface{}{
		"id": "001",
		"name": "Macbook",
		"price": "10USD",
		"image_url": "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
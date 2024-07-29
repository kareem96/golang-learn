package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id string `json:"id"` // add json as tag
	Name string `json:"name"`
	Price string `json:"price"`
	ImageUrl string `json:"image_url"`
}
func TestJsonTag(t *testing.T)  {
	product := Product{
		Id: "001",
		Name: "Macbook",
		Price: "10USD",
		ImageUrl: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
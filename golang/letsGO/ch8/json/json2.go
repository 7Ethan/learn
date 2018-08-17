package main

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Contact1 struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact_1"`
	Contact2 struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact_2"`
}

var JSON = `{
	"name" : "Gopher",
	"title" : "programmer",
	"contact_1" : {
		"home" : "shanghai",
		"cell" : "1271234143"
	},
	"contact_2" : {
		"home" : "beijingh",
		"cell" : "88888888"
	}
}`

func main() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c)
}

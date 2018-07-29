package main

import (
	"encoding/json"
	"fmt"
)

type _Contact struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Contact1 struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact1"`
	Contact2 struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact2"`
}

var _JSON = `{
	"name" : "Gopher",
	"title" : "programmer",
	"contact1" : {
		"home" : "shanghai",
		"cell" : "1271234143"
	},
	"contact2" : {
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

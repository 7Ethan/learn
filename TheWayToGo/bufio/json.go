package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"ahah", "jaja", "sadf"}
	wa := &Address{"dafa", "dasfa", "da"}
	vc := &VCard{"Sergey", "Jay", []*Address{wa, pa}, "none"}

	//js ,_ := json.Marshal(vc)
	//fmt.Printf("json format : %s",js)
	fmt.Println()
	file, _ := os.OpenFile("vacard.json", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		panic(err)
	}
}

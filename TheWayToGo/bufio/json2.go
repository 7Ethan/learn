package main

import "encoding/json"

func main() {
	//b  := []byte({"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]})
	var f interface{}
	err := json.Unmarshal(b, f)
	if err != nil {
		return
	}
}

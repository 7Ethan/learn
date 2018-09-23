package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}
	var status uint64
	err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status)
	// checkError(err)
	fmt.Println("Status value: ", status, err)
	/*
	   Status value:  200 <nil>
	*/
}

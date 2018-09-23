package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	var data = []byte(`{"status": 200}`)
	var result struct {
		Status uint64 `json:"status"`
	}

	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println(result)
	//{200}
}

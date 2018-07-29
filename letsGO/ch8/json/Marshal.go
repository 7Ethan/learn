package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	f, err := os.Create("jsonData.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := &Student{"zhang san", 18}
	encoder := json.NewEncoder(f)
	encoder.Encode(s)
}

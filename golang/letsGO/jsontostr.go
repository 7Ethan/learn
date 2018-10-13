package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var u User
	h := `{"name":"张三","age":15}`
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(u)
	}
}

type User struct {
	Name string `name`
	Age  int    `age`
}

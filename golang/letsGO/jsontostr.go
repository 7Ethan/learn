package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var u User
	h := `{
			{"name":"张三","age":15},
			{"name":"lisi","age":20}
		}`
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
}

type User struct {
	Name string `name`
	Age  int    `age`
}

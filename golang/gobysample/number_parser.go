package main

import (
	"fmt"
	"strconv"
)

func main(){
	f,_ := strconv.ParseFloat("12.345",64)
	fmt.Println(f)
}

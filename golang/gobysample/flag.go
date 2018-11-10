package main

import (
	"flag"
)

func main() {
	var str string
	flag.StringVar(&str, "s", "this is a string for flag", "str")
	flag.Parse()
}

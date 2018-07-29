package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type ParsrError struct {
	Index int
	Word  string
	Err   error
}

func (e *ParsrError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
	/*
		注意区分!:
			return fmt.Sprintf("pkg parse: error parsing %q as int", e)
	*/
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg : %v", r)
			}
		}
	}()
	fields := strings.Fields(input)
	numbers = field2numbers(fields)
	return
}

func field2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse ")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParsrError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return numbers
}

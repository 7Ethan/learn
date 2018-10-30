package main

import (
	"fmt"
	"reflect"
)

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func sliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func stringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func main() {
	a := []string{"datgagwer4taert", "sdfar", "saq23er"}
	b := []string{"datgagwer4taert", "sdfar", "saq23er"}
	fmt.Println(sliceEqual(a, b))
	fmt.Println(stringSliceReflectEqual(a, b))
	fmt.Println(sliceEqualBCE(a, b))
}

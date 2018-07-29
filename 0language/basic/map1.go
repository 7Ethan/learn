package language

import "fmt"

func main() {
	//m2 := make(map[int]string,100)
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m1)
	println(m1["v"])
	m := map[int]struct {
		name string
		age  int
	}{
		1:  {"haha", 10},
		20: {"hehe", 20},
	}
	println(m[20].age)
}

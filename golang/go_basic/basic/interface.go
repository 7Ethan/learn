package main

type Stringer interface {
	String() string
}

var b interface{}

func main() {

	switch s := b.(type) {
	case string:
		return s
	case Stringer:
		return s.String()
	}
}

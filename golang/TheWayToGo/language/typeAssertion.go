package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float64
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Println(t)
	}
	switch areaIntf.(type) {
	case *Square:
		fmt.Println("Square")
	case *Circle:
		fmt.Println("Circle")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("I don't know")
	}
}

func (sq *Square) Area() float64 {
	return float64(sq.side * sq.side)
}

func (c *Circle) Area() float64 {
	return float64(c.radius * c.radius * math.Pi)
}

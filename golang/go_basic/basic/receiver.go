package language

import "fmt"

type data struct {
	x int
}

func (receiver data) valueTest() {
	receiver.x = 100
}

func (receiver *data) pointerTest() {
	receiver.x = 200
}

func main() {
	d := data{}
	p := &d

	p.valueTest()
	fmt.Println(d)
	p.pointerTest()
	fmt.Println(d)

	d.valueTest()
	fmt.Println(d)
	d.pointerTest()
	fmt.Println(d)
}

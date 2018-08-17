package main

import "fmt"

//import "debug/dwarf"

type Node struct {
	le   *Node
	data []interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) Setdata(data []interface{}) error {
	n.data = data
	return nil
}

//var ai AbsIn

type myPrintInterface interface {
	printa()
}

type testInterface interface {
}

//func (t testInterface)printa()  {
//
//}
type test struct {
	a int
}

func (t test) printa() {
	fmt.Println("-------")
}

func f1(x testInterface) {
	x.(myPrintInterface).printa() //第一,传入的参数必须是interface类型,第二,这个是动态转换
}

func main() {
	var v test
	f1(v)
}

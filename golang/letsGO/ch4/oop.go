package main

import "fmt"

func main() {
	f := NewFruit()
	a := NewApple()
	f.DisplayName()
	a.DisplayName()
}

//定义一个函数类型
type FruitName func() string

type Fruit struct {
	GetFruitName FruitName
}

func (this *Fruit) DisplayName() {
	fmt.Println(this.GetFruitName()) //用一个函数签名
}

func (this *Fruit) GetName() string {
	return "水果"
}

func NewFruit() *Fruit {
	f := new(Fruit)
	f.GetFruitName = f.GetName
	return f
}

type Apple struct {
	Fruit
}

func (this *Apple) GetName() string {
	return "Apple"
}

func NewApple() *Apple {
	a := new(Apple)
	a.GetFruitName = a.GetName //注意要使用函数签名
	//cannot use a.GetName() (type string) as type FruitName in assignment
	return a
}

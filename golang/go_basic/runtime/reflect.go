package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) Say() {
	fmt.Println("My name is :", u.Name)
}

func (u *User) ChangeEmail(mail string) {
	u.Email = mail
}

func (u User) print() {
	fmt.Println(u.Name)
	fmt.Println(u.Email)
	fmt.Println(u.Age)
}

func main() {
	var u User = User{
		Name:  "Ethan",
		Age:   19,
		Email: "123@qq.com",
	}
	// fmt.Println(u)
	u.print()
	fmt.Println("----------------")
	u.ChangeEmail("hahahh@gmail.com")
	GetFieldAndMethod(u)
}

func GetFieldAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

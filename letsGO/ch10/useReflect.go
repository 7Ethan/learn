package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (this *Student) Printname() {
	fmt.Println(this.Name)
}

func (this *Student) GetAge() int {
	return this.Age
}

func main() {
	s := Student{"lilei", 18}
	rt := reflect.TypeOf(s)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		fmt.Println("----")
	}
	//fmt.Println("rt's type is : ",rt.Elem())
	fmt.Println(rt.PkgPath())
	//反射取所有字段
	fmt.Println(rt.Name(), "共有", rt.NumField(), "个字段")
	for i, j := 0, rt.NumField(); i < j; i++ {
		rtField := rt.Field(i)
		fmt.Println(rtField.Name)
	}
	/*因为我们的函数定义是在*Student 类型上的，所以这里转换为指针类型，否 则反射会取不到函数*/
	rt = reflect.PtrTo(rt)
	//反射取所有函数
	fmt.Println(rt.Name(), "共有", rt.NumMethod(), "个函数")
	for i, j := 0, rt.NumMethod(); i < j; i++ {
		mt := rt.Method(i)
		fmt.Println(mt.Name)     //输入参数的数量
		numIn := mt.Type.NumIn() //输出参数的数量
		numOut := mt.Type.NumOut()
		//输出输入参数
		if numIn > 0 {
			fmt.Println("\t 共", numIn, "个输入参数")
			for k := 0; k < numIn; k++ {
				in := mt.Type.In(k)
				fmt.Println("\t", in.Name(), "\t", in.Kind())
			}
		}
		//输出输出参数
		if numOut > 0 {
			fmt.Println("\t 共", numOut, "个输出参数")
			for k := 0; k < numOut; k++ {
				out := mt.Type.Out(k)
				fmt.Println("\t", out.Name(), "\t", out.Kind())
			}
		}
	}
}

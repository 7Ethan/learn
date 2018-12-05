/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 15:14
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//
//type student struct {
//	Name string
//	Age  int
//	}
//	func pase_student() {
//	m := make(map[string]*student)
//	stus := []student{
//		{Name: "zhou", Age: 24},
//		{Name: "li", Age: 23},
//		{Name: "wang", Age: 22},
//	}
//	for _, stu := range stus {
//		m[stu.Name] = &stu
//	}
//	}

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i-->: ", i) //闭包里引用了不作为参数传递进去的值,都是引用传递
			//当执行到这里打印的时候,for循环已经跑完,而且i++,所以会打印十个10
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

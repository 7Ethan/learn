/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 12:35
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */
package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "Excellent" {
		talk = "Nice"
	} else {
		talk = "Hi"
	}
	return
}

func main() {
	var p People = Student{}
	think := "Excellent"
	fmt.Println(p.Speak(think))
}

/*
两种方案:
一:	func (stu *Student)Speak(think string)(talk string){}
	main函数中 : var p People  = &Student{}

二:func (stu Student)Speak(think string)(talk string)
	main函数中var p People  = Student{} 或 var p People  = &Student{}
*/

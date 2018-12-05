/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 20:25
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
	p.ShowA()
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func (t *Teacher) ShowA() {
	fmt.Println("teacher--> showA")
}

func main() {
	t := Teacher{}
	t.ShowA()
}

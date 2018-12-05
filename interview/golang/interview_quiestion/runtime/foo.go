/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 20:08
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package main

import (
	"fmt"
	"time"
)

var qiut chan struct{}

func foo(id int) {
	fmt.Println(id)
	time.Sleep(time.Second)
	qiut <- struct{}{}
}
func main() {
	count := 1000
	qiut = make(chan struct{}, count)
	for i := 0; i < count; i++ {
		go foo(i)
	}

	for i := 0; i < count; i++ {
		<-qiut
	}
}

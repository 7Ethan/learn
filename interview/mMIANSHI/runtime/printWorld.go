/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 20:14
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package main

import (
	"fmt"
	_ "runtime"
)

func say(string string) {
	for i := 0; i < 5; i++ {
		fmt.Println(string)
	}
}

func main() {
	//runtime.GOMAXPROCS(1)	//如果开启这行代码则会什么都不打印,因为for占用了单核cpu的所有资源
	//gogland环境默认开启多线程
	go say("hello")
	for {
	}
}

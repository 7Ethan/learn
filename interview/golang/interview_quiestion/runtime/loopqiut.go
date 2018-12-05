/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 18:50
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package runtime

import (
	"fmt"
	"runtime"
	"time"
)

var qiut = make(chan struct{})

func loop() {
	for i := 0; i < 100; i++ {
		fmt.Print(" ", i)
		time.Sleep(10 * time.Millisecond)
	}
	qiut <- struct{}{}
}
func loop1() {
	for i := 0; i < 100; i++ {
		fmt.Print("-", i)
		time.Sleep(10 * time.Millisecond)
	}
	qiut <- struct{}{}
}

func main() {
	runtime.GOMAXPROCS(4)
	go loop1()
	go loop()
	for i := 0; i < 2; i++ {
		<-qiut
	}
}

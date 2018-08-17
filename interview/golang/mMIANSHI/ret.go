/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 20:38
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 10 //此处会将ret赋值为10,然后同时执行延迟函数defer,所以返回值是11
}
func main() {
	fmt.Println(f())
}

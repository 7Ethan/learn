/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 12:20
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */
package main

import "fmt"

func main() {
	s := make([]int, 3)
	fmt.Println(s)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  07/10/2017 13:09
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	p, err := filepath.Abs(filepath.Dir(os.Args[0])) //filepath.Abs-->change path to abslute Path
	if err != nil {
		log.Fatal("--->", err)
	}
	fmt.Println(p)
	fmt.Println("===========")
	fmt.Println(path())
}

func path() string {
	return filepath.Dir(os.Args[0])
}

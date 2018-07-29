package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	fmt.Println(path.Base("/usr/bin"))
	fmt.Println(path.Base(""))            //输出.
	fmt.Println(path.Base("C:\\Windows")) /*无法识别Windows下的路径分隔符，将会把 C:\\Windows 做为一个路径*/
	fmt.Println(path.Base(strings.Replace("C:\\Windows", "\\", "/", -1)))
	/*把\转换成/*/
}

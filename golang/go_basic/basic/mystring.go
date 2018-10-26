package main

import "fmt"

type Mystring string

func (m Mystring) String() string {
	//return fmt.Sprintf("Mystrinbg is %s",m)	//这里已经将String函数构造成了无穷递归
	return fmt.Sprintf("Mystrinbg is %s", string(m))
}
func main() {
	var str Mystring = "hahahah"
	str.String()
}

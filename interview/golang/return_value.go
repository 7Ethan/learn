package main

func main() {

}

func funcMui(x, y int) (sum int, error) {
	//在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。 
	//如果返回值有有多个返回值必须加上括号； 如果只有一个返回值并且有命名也需要加上括号
	return x + y, nil
}

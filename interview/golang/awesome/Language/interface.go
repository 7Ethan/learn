package Language

import "fmt"

type IA interface{}
type IB interface{ foo() }

type TA struct{}
type TB struct{}

func (TB) foo() {

}

//A := 100

func main() {
	var ia IA
	//var ta  = ia.(TA)
	var ta = ia
	var ib IB
	var tb = ib.(TB)

	//接口类型向普通类型是类型断言(运行期确定) 这个是很危险的
	//类型断言在编译期是没有任何保障的, 错误的代码也可以编译通过: 如var tb = ib.(TB)
	//panic: interface conversion: main.IB is nil, not main.TB

	fmt.Println(ta)
	fmt.Println(tb)
	fmt.Println()
}

//
//func x 转换为 y
//if x 是接口吗 ? {
//if 可以编译期转换吗 ? {
//这是类型转换 (接口之间可以隐式转换)
//// 这里也可以用类型断言, 如果编译期不优化的效率可能低一些
//} else {
//这是类型断言 (运行时也可能会失败)
//}
//} else {
//if 可以编译期转换 ? {
//这是类型转换 (显示转换, 必须成功, 但可能会丢失数据)
//} else {
//禁止!
//}
//}
//}

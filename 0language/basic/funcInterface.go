package language

//让函数直接实现接口
type Tester interface {
	do()
}

type FuncDo func()

func (self FuncDo) do() {
	self()
}
func main() {
	var t Tester = FuncDo(func() {
		println("--------")
	})
	t.do()
}

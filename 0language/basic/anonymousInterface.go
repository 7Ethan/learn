package language

import (
	"fmt"
	"reflect"
)

type tester struct {
	s interface {
		String() string
	}
}

type user struct {
	id   int
	name string
}

func (self *user) String() string {
	return fmt.Sprintf("User %d ,%s", self.id, self.name)
}
func main() {
	//t := user{id : 001,}
	t := tester{&user{id: 001, name: "Sergey"}}
	//switch value := t.(type) {
	//case string:
	//	fmt.Println("str",value)
	//case int:
	//	fmt.Println("int",value)
	//default:
	//	fmt.Println("don't know")
	//}
	//a := "......"
	//if v,ok := a.(string);ok {
	//	fmt.Println(v)
	//}

	a := reflect.TypeOf(t)
	b := reflect.ValueOf(t)
	fmt.Println(a.Field(0).Type)

	fmt.Println(a.Kind())
	fmt.Printf("%T", b)
	fmt.Println(b.Kind())

	fmt.Println(t.s.String())
}

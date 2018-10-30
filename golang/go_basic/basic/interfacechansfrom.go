package language

import "fmt"

type Stringer interface {
	String() string
}

type Printer interface {
	String() string
	Print()
}

type user struct {
	id   int
	name string
}

func (self *user) String() string {
	return fmt.Sprintf("user : %d ,name : %s", self.id, self.name)
}

func (self *user) Print() {
	fmt.Println(self.String())
}

func main() {
	//var 	u = user{id:002,name:"Jay"}
	////var _ user.Printer  = (*Data)(nil)
	//var _ Stringer  = (u)(nil)
	var o Printer = &user{id: 001, name: "Sergey"}
	var s Stringer = o
	fmt.Println(s.String())
	//fmt.Println(s.)
}

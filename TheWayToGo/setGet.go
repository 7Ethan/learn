package TheWayToGo

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(firstName string) error {
	p.firstName = firstName
	return nil
}

//func main() {
//	var p Person = Person{"Sergey","Jay"}
//	fmt.Printf("%#v",p)
//}

func (p *Person) String() string {
	//return "Person : {" + "\\n" + "first : " + p.firstName +
	//	",\\n" + "lastName : " + p.lastName + "\\n}"
	return fmt.Sprintf("%v", p)
}

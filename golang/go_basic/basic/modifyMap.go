package language

import "fmt"

type file struct {
	name string "tag"
	size int
	attr struct {
		pern int
		ower int
	}
}

func main() {

	f := file{
		name: "hahh",
		size: 100,
		//	attr:struct{
		//		pern:0,
		//		ower:0,
		//	}
	}
	f.attr.pern = 0
	f.attr.ower = 0

	type user struct {
		name string
	}

	m := map[string]user{
		"1": {"user"},
	}

	fmt.Println(m)
	u := m["1"]
	u.name = "newuser"
	m["1"] = u
	fmt.Println(m)

}

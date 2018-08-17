package Language

type S struct {
	name string
}

func main() {
	m := map[string]*S{"x": &S{"one"}}
	m["x"].name = "two"

	if v, ok := m["x"]; ok {
		println(v.name)
	}
}

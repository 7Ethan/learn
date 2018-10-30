package main

type errDemo struct {
}
type ERR struct {
}

func (e *errDemo) Error() string {
	return "errDemo erroe!"
}

func (e *ERR) Error() string {
	return "errDemo erroe!"
}

type ERROR struct {
	ERR
	errDemo
}

func main() {
	//var err error = errDemo{}
	//var err error = (*errDemo)(nil)
	//println(err == nil)

	//---------

	var e ERROR = ERROR{}
	e.Error()

}

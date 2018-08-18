package language

import "errors"

var ErrDivByZero = errors.New("Division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}
func main() {
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		//println(err)
		panic(err)
	}
}

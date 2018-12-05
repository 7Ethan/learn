package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(^uint32(0)>>1, "----")
	fmt.Println(Reverse(-2147483648))
	fmt.Println(Reverse(-2147483647))
	fmt.Println(Reverse(2147483647))
	fmt.Println(Reverse(-21474))
	fmt.Println(Reverse(21474))
	fmt.Println(Reverse(0))
	//fmt.Println(Reverse2(6490542445))

	//fmt.Println(Reverse2(-2147483648))
}

func Reverse(x int32) int32 {
	max := int32(^uint32(0) >> 1)
	if -(x + 1) >= max {
		return 0
	}
	var c, a int32
	var i int
	if x < 0 {
		c = -x
	} else {
		c = x
	}
	for {
		a = c % 10
		c = c / 10
		i = i*10 + int(a)
		if i > int(max) {
			return 0
		}
		if c == 0 {
			break
		}
	}
	if x < 0 {
		return -int32(i)
	}
	return int32(i)
}

func Reverse2(x int64) int64 {
	var ret int64
	sign := 0
	if x < 0 {
		x = x * -1
		sign = 1
	}

	len, tmp := 0, x
	for {
		if tmp < 0 || tmp == 0 {
			break
		}
		len = len + 1
		tmp = tmp / 10
	}

	for i := len; i > 0; i-- {
		j := len - i
		if j != 0 {
			ret = ret*10 + x%int64(math.Pow10(j+1))/int64(math.Pow10(j))
		} else {
			ret = ret + x%10
		}
		fmt.Println(i, ret)
	}
	if sign != 0 {
		ret = ret * -1
	}
	return ret
}

package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := myPow(2.00000, 10)
	fmt.Println(res)
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.00000
	}
	if n == 1 {
		return x
	}
	if x == 1.00000 {
		return x
	}
	if x == -1.00000 {
		if n%2 == 0 {
			return 1.00000
		} else {
			return -1.00000
		}
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	res := 1.00000
	for n != 0 {
		// it should be on the before x *= x
		if n&1 == 1 {
			res *= x
		}
		// get x value on the next loop
		x *= x
		n = (n >> 1)
	}
	return res
}

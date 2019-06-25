package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := plusOne([]int{9})
	fmt.Println(res)
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	// res := make([]int, len(digits)+1)
	length := len(digits)
	digits[length-1] = digits[length-1] + 1
	digits = append(digits, 0)
	for i := length - 1; i >= 0; i-- {
		digits[i+1] = digits[i] % 10
		if i == 0 {
			if digits[i] >= 10 {
				digits[i] = digits[i] / 10
			} else {
				digits[i] = 0
			}
		} else {
			digits[i-1] += digits[i] / 10
		}

	}
	if digits[0] != 0 {
		return digits
	}
	return digits[1:]
}

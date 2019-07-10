package main

import (
	"fmt"
	// "math"
)

func main() {
	/*
		"968018661806000118986811000908199810896"
		"96"
		"88"
		"101"
	*/

	res := isStrobogrammatic("101")
	fmt.Println(res)
}

func isStrobogrammatic(num string) bool {
	// oriNum := num
	if num == "0" {
		return true
	}
	rotated := make([]byte, 0)
	for i := len(num) - 1; i >= 0; i-- {
		digit := num[i]
		if digit == '6' {
			rotated = append(rotated, '9')
			continue
		}
		if digit == '9' {
			rotated = append(rotated, '6')
			continue
		}
		if digit == '8' {
			rotated = append(rotated, '8')
			continue
		}
		if digit == '1' {
			rotated = append(rotated, '1')
			continue
		}
		if digit == '0' {
			rotated = append(rotated, '0')
			continue
		}
		return false
	}

	// fmt.Println(string(rotated))
	// fmt.Println(num)
	if string(rotated) == num {
		return true
	}
	return false
}

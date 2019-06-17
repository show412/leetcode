package main

import (
	"fmt"
)

func main() {
	res := isPalindrome(23532)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	originNum := x
	reverseNum := 0
	for originNum != 0 {
		mod := originNum % 10
		originNum = originNum / 10
		reverseNum = reverseNum*10 + mod
	}
	if reverseNum == x {
		return true
	}
	return false
}

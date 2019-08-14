package main

import (
	"fmt"
	"strings"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := isPalindrome("0P")
	fmt.Println(res)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1
	for start != end && start <= end {
		for start < len(s) && checkCase(s[start]) == false {
			start++
		}
		for end >= 0 && checkCase(s[end]) == false {
			end--
		}
		if start <= end && s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func checkCase(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}

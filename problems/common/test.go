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
	res := repeatedStringMatch("aaaaaaaaaaaaaaaaaaaaaab", "ba")
	fmt.Println(res)
}

func repeatedStringMatch(A string, B string) int {
	S := A
	res := 1
	for ; len(S) < len(B); res++ {
		S = S + A
	}
	if strings.Index(S, B) >= 0 {
		return res
	}
	if strings.Index(S+A, B) >= 0 {
		return res + 1
	}
	return -1
}

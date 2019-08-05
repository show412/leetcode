package main

import (
	"fmt"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
		return res
	}
	for c := 0; c < n; c++ {
		for _, left := range generateParenthesis(c) {
			for _, right := range generateParenthesis(n - 1 - c) {
				res = append(res, "("+left+")"+right)
			}
		}
	}
	return res
}

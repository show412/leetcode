package main

import (
	"fmt"
)

/*
test case:
[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3
10

[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2
6
*/
func main() {
	res := longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	// 998001
	fmt.Println(res)
}

func longestOnes(A []int, K int) int {
	if K == len(A) {
		return len(A)
	}
	s, e := 0, 0
	res := 0
	n := 0
	for s <= e && e < len(A) {
		if A[e] == 0 {
			n++
		}
		for n > K {
			if A[s] == 0 {
				n--
			}
			s++
		}
		e++
		res = max(res, e-s)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

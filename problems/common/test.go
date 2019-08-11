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
	res := intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}})
	fmt.Println(res)
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	i := 0
	j := 0
	m := len(A)
	n := len(B)
	res := make([][]int, 0)
	for i < m && j < n {
		// it's wrong to split interval
		if B[j][0] >= A[i][0] && B[j][0] <= A[i][1] {
			res = append(res, []int{max(A[i][0], B[j][0]), min(A[i][1], B[j][1])})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

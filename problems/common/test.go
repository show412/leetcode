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
	res := sortedSquares([]int{-1, 2, 2})
	// 998001
	fmt.Println(res)
}

func sortedSquares(A []int) []int {
	res := make([]int, 0)
	if A[len(A)-1] <= 0 {
		for i := len(A) - 1; i >= 0; i-- {
			res = append(res, A[i]*A[i])
		}
		return res
	}
	if A[0] >= 0 {
		for i := 0; i < len(A); i++ {
			res = append(res, A[i]*A[i])
		}
		return res
	}
	p1 := 0
	p2 := 0
	for i := 0; i < len(A); i++ {
		if A[i] >= 0 && A[i-1] <= 0 {
			p1 = i - 1
			p2 = i
			break
		}
	}
	for p1 >= 0 && p2 <= len(A)-1 {
		if A[p1]*A[p1] < A[p2]*A[p2] {
			res = append(res, A[p1]*A[p1])
			p1--
		} else {
			res = append(res, A[p2]*A[p2])
			p2++
		}
	}
	for p1 >= 0 {

		res = append(res, A[p1]*A[p1])
		p1--

	}
	for p2 <= len(A)-1 {

		res = append(res, A[p2]*A[p2])
		p2++

	}
	return res
}

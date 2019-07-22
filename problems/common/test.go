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
	res := minDominoRotations([]int{2, 1, 1, 3, 2, 1, 2, 2, 1}, []int{3, 2, 3, 1, 3, 2, 3, 3, 2})
	fmt.Println(res)
}

func minDominoRotations(A []int, B []int) int {
	if len(A) == 0 {
		return -1
	}
	if len(A) == 1 {
		return 0
	}
	resA := 0
	resAB := 0
	flagA := false
	// A[0] is the standard
	for i := 1; i < len(A); i++ {
		if A[i] == A[0] {
			flagA = true
			resAB++
			continue
		} else if B[i] == A[0] {
			resA++
		} else {
			flagA = false
			break
		}
		flagA = true
	}
	fmt.Println(resA)
	fmt.Println(resAB)
	ra := min(resA, resAB)
	// B[0] is the standard
	resB := 0
	resBA := 0
	flagB := false
	// A[0] is the standard
	for i := 1; i < len(B); i++ {
		if B[i] == B[0] {
			flagB = true
			resBA++
			continue
		} else if A[i] == B[0] {
			resB++
		} else {
			flagB = false
			break
		}
		flagB = true
	}
	fmt.Println(resB)
	// fmt.Println(len(B))
	rb := min(resB, resBA)
	if flagA == false && flagB == false {
		return -1
	}
	if flagA == true && flagB == true {
		return min(ra, rb)
	}
	if flagA == true {
		return ra
	}
	return rb
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

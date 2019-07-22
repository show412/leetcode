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
	for i := 0; i < len(A); i++ {
		if A[i] != A[0] && B[i] != A[0] {
			flagA = false
			break
		} else if B[i] != A[0] {
			resAB++
			// continue
		} else if A[i] != A[0] {
			resA++
		}
		flagA = true
	}

	resA = min(resA, resAB)

	// B[0] is the standard
	resB := 0
	resBA := 0
	flagB := false
	// A[0] is the standard
	for i := 0; i < len(B); i++ {
		if A[i] != B[0] && B[i] != B[0] {
			flagB = false
			break
		} else if B[i] != B[0] {
			resBA++
		} else if A[i] != B[0] {
			resB++
		}
		flagB = true
	}

	resB = min(resB, resBA)
	if flagA == false && flagB == false {
		return -1
	}
	if flagA == true && flagB == true {
		return min(resA, resB)
	}
	if flagA == true {
		return resA
	}
	return resB
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

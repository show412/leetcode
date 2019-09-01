package main

import (
	"fmt"
	// "math"
)

/*
test case:
[0,0]
-1
true

[0,1,0]
0
false

[23, 2, 6, 4, 7],  k=0
false

[5,0,0]
0
true

{23, 2, 6, 4, 7}, 6
true

[23, 2, 4, 6, 7], 6
true
*/
func main() {
	res := isMonotonic([]int{1, 1, 1})
	fmt.Println(res)
}

func isMonotonic(A []int) bool {
	asc := true
	dec := true
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			asc = false
			break
		}
	}
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			dec = false
			break
		}
	}
	if asc == false && dec == false {
		return false
	}
	return true
}

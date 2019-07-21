package main

import (
	"fmt"
	"math"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := maxDistToClosest([]int{1, 0, 0, 0})
	fmt.Println(res)
}

func maxDistToClosest(seats []int) int {
	Forward := make([]int, len(seats))
	Afterward := make([]int, len(seats))
	curF := -1
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			curF = i
			Forward[i] = 0
			continue
		}
		if curF == -1 {
			Forward[i] = math.MaxInt32
		} else {
			Forward[i] = i - curF
		}
	}
	curA := -1
	for i := len(seats) - 1; i >= 0; i-- {
		if seats[i] == 1 {
			curA = i
			Afterward[i] = 0
			continue
		}
		if curA == -1 {
			Afterward[i] = math.MaxInt32
		} else {
			Afterward[i] = curA - i
		}
	}
	res := 0
	for i := 0; i < len(seats); i++ {
		res = max(res, min(Forward[i], Afterward[i]))
	}
	return res
}

func max(a int, b int) int {
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

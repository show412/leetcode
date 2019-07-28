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
	res := shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3)
	fmt.Println(res)
}

func shipWithinDays(weights []int, D int) int {
	// res := 0
	// frontSum := make([]int, len(weights))
	// frontSum[0] = weights[0]
	start := 0
	end := 0
	for i := 0; i < len(weights); i++ {
		start = max(start, weights[i])
		end += weights[i]
	}
	if D == 1 {
		return end
	}
	for start < end {
		mid := start + (end-start)/2
		if checkCapacity(weights, mid, D) > 0 {
			start = mid + 1
		} else {
			end = mid
		}

	}
	return start
}

func checkCapacity(weights []int, capacity int, D int) int {
	count := 1
	sum := 0
	// fmt.Println(capacity)
	for i := 0; i < len(weights); i++ {
		if (sum + weights[i]) > capacity {
			count++
			sum = weights[i]
		} else {
			sum += weights[i]
		}
	}
	if count > D {
		return 1
	} else if count < D {
		return -1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

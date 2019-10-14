package main

import (
	"fmt"
)

/*
test case:
[[1,1,4],[9,4,9],[9,1,9],[2,3,5],[4,1,5],[10,4,5]]
33
false

*/

func main() {
	res := longestWPI([]int{})
	// 998001
	fmt.Println(res)
}
func longestWPI(hours []int) int {
	res := 0
	// record the score is negative and index
	m := make(map[int]int, 0)
	score := 0
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			score += 1
		} else {
			score -= 1
		}
		if score > 0 {
			res++
		} else {
			if v, ok := m[score-1]; ok {
				res = max(res, i-v+1)
			}
			m[score] = i
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

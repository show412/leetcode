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
	res := findMaximumXOR([]int{3, 10, 5, 25, 2, 8})
	// 998001
	fmt.Println(res)
}
func findMaximumXOR(nums []int) int {
	res := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask = mask | (1 << uint(i))
		m := make(map[int]bool, 0)
		for _, num := range nums {
			m[mask&num] = true
		}
		temp := res | (1 << uint(i))
		for pre, _ := range m {
			if m[pre^temp] == true {
				res = temp
				break
			}
		}
	}
	return res
}

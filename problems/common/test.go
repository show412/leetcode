package main

import (
	"fmt"
)

/*
test case:
"a, a, a, a, b,b,b,c, c"
["a"]
"b"

paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
*/
func main() {
	res := longestArithSeqLength([]int{83, 20, 17, 43, 52, 78, 68, 45})
	// res := canFinish(4, [][]int{[]int{0, 1}})
	fmt.Println(res)
}

// type m map[int]int

func longestArithSeqLength(A []int) int {
	if len(A) == 2 {
		return 2
	}
	max := 2
	// 在第几个数差值有多少个 这个数据结构是这道题的关键
	dp := make([]map[int]int, len(A))
	for i := 0; i < len(A); i++ {
		x := A[i]
		dp[i] = make(map[int]int, 0)
		for j := 0; j < i; j++ {
			y := A[j]
			d := x - y
			if _, ok := dp[j][d]; !ok {
				dp[j][d] = 1
			}
			dp[i][d] = dp[j][d] + 1
			if dp[i][d] > max {
				max = dp[i][d]
			}
		}
	}
	return max
}

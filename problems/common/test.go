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
	res := mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})
	// 998001
	fmt.Println(res)
}

func mincostTickets(days []int, costs []int) int {
	if len(days) == 1 {
		return costs[0]
	}
	memo := make([]int, 51)
	dayset := make(map[int]bool, 0)
	for _, v := range days {
		dayset[v] = true
	}
	return dp(1, &memo, dayset, costs)
}

func dp(i int, memo *[]int, dayset map[int]bool, costs []int) int {
	if i > 50 {
		return 0
	}
	if (*memo)[i] != 0 {
		return (*memo)[i]
	}
	var ans int
	if _, ok := dayset[i]; ok {
		ans = min(dp(i+1, memo, dayset, costs)+costs[0], dp(i+7, memo, dayset, costs)+costs[1])
		ans = min(ans, dp(i+30, memo, dayset, costs)+costs[2])
	} else {
		ans = dp(i+1, memo, dayset, costs)
	}
	(*memo)[i] = ans
	fmt.Println("****")
	fmt.Println(i)
	fmt.Println(*memo)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

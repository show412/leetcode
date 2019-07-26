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
	res := assignBikes([][]int{[]int{0, 0}, []int{2, 1}}, [][]int{[]int{1, 2}, []int{3, 3}})
	fmt.Println(res)
}

func assignBikes(workers [][]int, bikes [][]int) int {
	res := math.MaxInt64
	cur := 0
	visit := make(map[int]bool, len(bikes))
	dfs(visit, workers, 0, bikes, cur, &res)
	return res
}

func dfs(visit map[int]bool, workers [][]int, start int, bikes [][]int, cur int, res *int) {
	if start >= len(workers) {
		*res = min(cur, *res)
		return
	}
	if cur > *res {
		return
	}
	for i := 0; i < len(bikes); i++ {
		if visit[i] == true {
			continue
		}
		visit[i] = true
		dfs(visit, workers, start+1, bikes, cur+dis(workers[start], bikes[i]), res)
		visit[i] = false
	}
}

func dis(worker []int, bike []int) int {
	return int(math.Abs(float64(worker[0]-bike[0]))) + int(math.Abs(float64(worker[1]-bike[1])))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

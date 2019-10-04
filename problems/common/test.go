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
	// [][]int{[]int{0, 1}, []int{0, 4}, []int{1, 4}, []int{2, 3}}
	// []int{0, 1}, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 4}
	// [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 4}}
	res := validTree(5, [][]int{[]int{0, 1}, []int{0, 4}, []int{1, 4}, []int{2, 3}})
	// 998001
	fmt.Println(res)
}

func validTree(n int, edges [][]int) bool {
	fa := make([]int, n)
	for i := 0; i < len(fa); i++ {
		fa[i] = i
	}
	for _, edge := range edges {
		unity(edge[0], edge[1], fa)
	}
	fmt.Println(fa)
	last := fa[0]
	for _, v := range fa {
		if v != last {
			return false
		}
	}
	return len(edges) == n-1
}

func unity(x int, y int, fa []int) {
	x = find(x, fa)
	y = find(y, fa)
	fa[y] = x
	return
}

func find(x int, fa []int) int {
	if fa[x] == x {
		return x
	} else {
		fa[x] = find(fa[x], fa)
	}
	return fa[x]
}

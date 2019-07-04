package main

import (
	"fmt"
	// "math"
)

func main() {

	res := removeStones([][]int{[]int{0, 0}, []int{0, 1}, []int{1, 0}, []int{1, 2}, []int{2, 1}, []int{2, 2}})
	fmt.Println(res)
}

func removeStones(stones [][]int) int {
	fa := make([]int, 1000)
	for i := 0; i < len(stones); i++ {
		fa[i] = i
	}

	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				unity(i, j, fa)
			}
		}

	}
	set := 0
	for i := 0; i < len(stones); i++ {
		if fa[i] == i {
			set++
		}
	}
	return len(stones) - set
}

func unity(x int, y int, fa []int) []int {
	x = find(x, fa)
	y = find(y, fa)
	fa[x] = y
	return fa
}

func find(x int, fa []int) int {
	if fa[x] == x {
		return x
	} else {
		fa[x] = find(fa[x], fa)
	}
	return fa[x]
}

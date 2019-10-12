package main

import (
	"fmt"
	"sort"
)

/*
test case:
[[1,1,4],[9,4,9],[9,1,9],[2,3,5],[4,1,5],[10,4,5]]
33
false

*/

func main() {
	res := carPooling([][]int{[]int{1, 1, 4}, []int{9, 4, 9}, []int{9, 1, 9}, []int{2, 3, 5}, []int{4, 1, 5}, []int{10, 4, 5}}, 33)
	// 998001
	fmt.Println(res)
}
func carPooling(trips [][]int, capacity int) bool {
	if len(trips) == 0 {
		return true
	}
	if len(trips) == 1 {
		if trips[0][0] <= capacity {
			return true
		}
		return false
	}
	m := make(map[int]int, 0)
	location := make([]int, 0)
	maxP := 0
	for i := 0; i < len(trips); i++ {
		p := trips[i][0]
		s := trips[i][1]
		e := trips[i][2]
		m[s] += p
		m[e] -= p
		location = append(location, s)
		location = append(location, e)
	}
	sort.Ints(location)
	for i := 0; i < len(location); i++ {
		maxP += m[location[i]]
		if maxP > capacity {
			return false
		}
	}
	return true
}

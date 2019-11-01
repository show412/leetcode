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
	res := canFinish(2, [][]int{[]int{0, 1}})
	// 998001
	fmt.Println(res)
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make(map[int][]int, 0)
	indgree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		s := prerequisites[i][0]
		e := prerequisites[i][1]
		edges[e] = append(edges[e], s)
		indgree[s]++
	}
	queue := make([]int, 0)
	fmt.Println(indgree)

	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			queue = append(queue, i)
		}
	}
	fmt.Println(queue)
	i := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		i++
		for _, edge := range edges[node] {
			if indgree[edge] > 0 {
				indgree[edge]--
			}
			if indgree[edge] == 0 {
				queue = append(queue, edge)
			}
		}
	}
	if i == numCourses {
		return true
	}
	return false
}

package main

import (
	"fmt"
	"sort"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
	fmt.Println(res)
}

func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 {
		return people
	}
	res := make([][]int, len(people))
	// sort.Sort should use the defined type
	sort.Sort(peopleSlice(people))
	fmt.Println(people)
	for i := 0; i < len(people); i++ {
		start := 0
		for j := 0; j < len(res); j++ {
			if res[j] != nil {
				continue
			}
			if start == people[i][1] {
				res[j] = people[i]
			}
			fmt.Println(res)
			start++
		}
	}
	return res
}

type peopleSlice [][]int

func (a peopleSlice) Len() int {
	return len(a)
}
func (a peopleSlice) Less(i, j int) bool {
	return a[i][0] < a[j][0] || (a[i][0] == a[j][0] && a[i][1] > a[j][1])
}
func (a peopleSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

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
	res := customSortString("cba", "abcd")
	// 998001
	fmt.Println(res)
}
func customSortString(S string, T string) string {
	orderMap := make(map[byte]int, 0)
	orderArray := make([]int, len(S))
	nonArray := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		orderMap[S[i]] = i
	}
	for i := 0; i < len(T); i++ {
		if v, ok := orderMap[T[i]]; ok {
			orderArray[v]++
		}
		if _, ok := orderMap[T[i]]; !ok {
			nonArray = append(nonArray, T[i])
		}

	}
	res := make([]byte, 0)
	for i := 0; i < len(orderArray); i++ {
		for orderArray[i] > 0 {
			res = append(res, S[i])
			orderArray[i]--
		}
	}
	res = append(res, nonArray...)
	return string(res)
}

package main

import (
	"fmt"
)

func main() {
	res := numJewelsInStones("z", "ZZ")
	fmt.Println(res)
}

func numJewelsInStones(J string, S string) int {
	arrJ := []rune(J)
	arrS := []rune(S)
	mapS := make(map[rune]bool, len(arrJ))
	cnt := 0
	for i := 0; i < len(arrJ); i++ {
		mapS[arrJ[i]] = true
	}
	for j := 0; j < len(arrS); j++ {
		if _, ok := mapS[arrS[j]]; ok {
			cnt++
		}
	}
	return cnt
}

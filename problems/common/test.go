package main

import (
	"fmt"
	// "math"
)

/*
test case:
["fxasxpc","dfbdrifhp","nwzgs","cmwqriv","ebulyfyve","miracx","sxckdwzv","dtijzluhts","wwbmnge","qmjwymmyox"]
"zkgwaverfimqxbnctdplsjyohu"
*/
func main() {
	res := isAlienSorted([]string{"fxasxpc", "dfbdrifhp", "nwzgs", "cmwqriv", "ebulyfyve", "miracx", "sxckdwzv", "dtijzluhts", "wwbmnge", "qmjwymmyox"}, "zkgwaverfimqxbnctdplsjyohu")
	fmt.Println(res)
}

func isAlienSorted(words []string, order string) bool {
	lexi := make(map[byte]int, 0)
	for k, v := range order {
		if _, ok := lexi[byte(v)]; !ok {
			lexi[byte(v)] = k
		}
	}
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]
		start := 0
		for start < len(word1) {
			if start >= len(word2) {
				return false
			}
			if lexi[word1[start]] < lexi[word2[start]] {
				break
			}
			if lexi[word1[start]] > lexi[word2[start]] {
				return false
			}
			start++
		}
	}
	return true
}

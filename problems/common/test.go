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
	res := findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb")
	// 998001
	fmt.Println(res)
}
func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		word := words[i]
		if match(word, pattern) == true {
			res = append(res, word)
		}
	}
	return res
}

func match(word string, pattern string) bool {
	m1 := make(map[byte]byte, 0)
	m2 := make(map[byte]byte, 0)
	for i := 0; i < len(word); i++ {
		if _, ok1 := m1[word[i]]; !ok1 {
			m1[word[i]] = pattern[i]
		}
		if _, ok2 := m2[pattern[i]]; !ok2 {
			m2[pattern[i]] = word[i]
		}
		if m2[pattern[i]] != word[i] || m1[word[i]] != pattern[i] {
			return false
		}
	}
	return true
}

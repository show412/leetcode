package main

import (
	"fmt"
	// "math"
	// "sort"
)

func main() {
	// [1,1],[1,3],[3,1],[3,3],[4,1],[4,3]
	// [3,2],[0,0],[3,3],[3,4],[4,4],[2,1],[4,3],[1,0],[4,1],[0,2]
	// a := [][]int{[]int{3, 2}, []int{0, 0}, []int{3, 3}, []int{3, 4}, []int{4, 4}, []int{2, 1}, []int{4, 3}, []int{1, 0}, []int{4, 1}, []int{0, 2}}
	b := getAns("aaaasds", "gkg")
	fmt.Println(b)
}

func getAns(s string, word string) string {
	// Write a code here
	if len(s) == 0 || len(word) == 0 || len(s) < len(word) {
		return "no"
	}
	slen := len(s)
	wlen := len(word)
	// Give `s="abcabcd"`, `word="xyzxyz"`, return `yes`
	// "jhdsgadhgsahdjgsahjdgashjdgshaj", "aaaaa"
	for i := 0; i < (slen - wlen + 1); i++ {
		j := 0
		m1 := make(map[string]string)
		m2 := make(map[string]string)
		for j < wlen {
			v1, ok1 := m1[string(s[i+j])]
			v2, ok2 := m2[string(word[j])]
			if !ok1 && !ok2 {
				m1[string(s[i+j])] = string(word[j])
				m2[string(word[j])] = string(s[i+j])
				j++
			} else if ok1 && ok2 && v1 == string(word[j]) && v2 == string(s[i+j]) {
				j++
			} else {
				break
			}
			if j == wlen {
				// fmt.Println(m1)
				// fmt.Println(m2)
				return "yes"
			}
		}
		i++
	}
	return "no"
}

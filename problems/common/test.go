package main

import (
	"fmt"
	"sort"
	// "math"
)

func main() {
	/*
		"abcd"
		[0, 2]
		["a", "cd"]
		["eee", "ffff"]
		expected: "eeebffff"
	*/

	/*
	   "vmokgggqzp"
	   [3,5,1]
	   ["kg","ggq","mo"]
	   ["s","so","bfr"]
	   expected: "vbfrssozp"
	*/

	res := findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"})
	fmt.Println(res)
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	res := make([]byte, 0)
	start := 0
	sourcesMap := make(map[int]string, len(sources))
	targetsMap := make(map[int]string, len(targets))
	for i := 0; i < len(indexes); i++ {
		sourcesMap[indexes[i]] = sources[i]
		targetsMap[indexes[i]] = targets[i]
	}
	sort.Ints(indexes)
	// fmt.Println(S[0:1] == sources[0])
	// fmt.Println(res = append(res, []byte(targets[0])...))

	for i := 0; i < len(S); i++ {
		if start < len(indexes) && i == indexes[start] {
			// fmt.Println("*****")
			// fmt.Println(string(S[i:(i + len(sources))]))
			// fmt.Println(sources[start])
			if string(S[i:(i+len(sourcesMap[indexes[start]]))]) == sourcesMap[indexes[start]] {
				res = append(res, []byte(targetsMap[indexes[start]])...)

				// fmt.Println(res)
				i = i + len(sourcesMap[indexes[start]]) - 1
			} else {
				res = append(res, S[i])
			}
			start++
		} else {
			res = append(res, S[i])
		}
	}
	return string(res)
}

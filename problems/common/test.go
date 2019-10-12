package main

import (
	"fmt"
)

/*
test case:
"adc"
"dcda"
true

"hello"
"ooolleoooleh"
false

"abc"
"bbbca"
true

"abcdxabcde"
"abcdeabcdx"
true

https://leetcode.com/submissions/detail/267059201/testcase/
false

https://leetcode.com/submissions/detail/267059572/testcase/
true
*/

func main() {
	res := lengthOfLongestSubstring("tmmzuxt")
	// 998001
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	l := 0
	r := 0
	m := make(map[byte]int, 0)
	for r < len(s) {
		if _, ok := m[s[r]]; ok {
			l = max(l, m[s[r]]+1)
		}
		m[s[r]] = r
		res = max(res, r-l+1)
		r++
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

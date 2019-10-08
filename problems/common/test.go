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
	res := findAnagrams("cbaebabacd", "abc")
	// 998001
	fmt.Println(res)
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	sa := make([]int, 26)
	pa := make([]int, 26)
	res := make([]int, 0)
	for i := 0; i < len(p); i++ {
		pa[p[i]-'a']++
	}
	left := 0
	right := 0
	for left < len(s) && right < len(s) {
		sa[s[right]-'a']++
		for left <= right && sa[s[right]-'a'] > pa[s[right]-'a'] {
			sa[s[left]-'a']--
			left++
		}
		if right-left+1 == len(p) {
			res = append(res, left)
			sa[s[left]-'a']--
			left++
		}
		right++
	}
	return res
}

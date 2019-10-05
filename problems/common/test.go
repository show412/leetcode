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
	res := checkInclusion("ab", "eidboaoo")
	// 998001
	fmt.Println(res)
}
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	arr1 := make([]int, 26)
	arr2 := make([]int, 26)
	// init the arr1 and arr2 at the same time to reduce the time
	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if match(arr1, arr2) {
			return true
		}
		arr2[s2[i+len(s1)]-'a']++
		arr2[s2[i]-'a']--
	}
	return match(arr1, arr2)
}

func match(arr1 []int, arr2 []int) bool {
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/permutation-in-string/
/*
Given two strings s1 and s2, write a function to return true
if s2 contains the permutation of s1.
In other words, one of the first string's permutations
is the substring of the second string.



Example 1:

Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input:s1= "ab" s2 = "eidboaoo"
Output: False


Note:

The input strings only contain lower case letters.
The length of both given strings is in range [1, 10,000].
*/
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

https://leetcode.com/submissions/detail/267068953/testcase/
true
*/
// this is a slide window
// use array to store character
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

// the solution is time limit exceeded
// the cost is on the s2 map initialize in second for loop
func checkInclusion(s1 string, s2 string) bool {
	ms1 := make(map[byte]int, 0)
	nums1 := 0
	for i := 0; i < len(s1); i++ {
		ms1[s1[i]] += 1
	}
	nums1 = len(ms1)
	l := 0
	for l <= (len(s2) - len(s1)) {
		ms2 := make(map[byte]int, 0)
		formNum := 0
		for ms1[s2[l]] == 0 && l <= (len(s2)-len(s1)) {
			l++
		}
		r := l
		// if (len(s2) - r) < len(s1) {
		// 	return false
		// }
		for (r-l) <= len(s1) && ms1[s2[r]] > 0 && ms2[s2[r]] <= ms1[s2[r]] {
			ms2[s2[r]]++
			if ms2[s2[r]] == ms1[s2[r]] {
				formNum++
			}

			// if _, ok := ms1[s2[r]]; !ok {
			// r++
			// ms2 = make(map[byte]int, 0)
			// formNum = 0
			// 	continue
			// }
			if ms2[s2[r]] > ms1[s2[r]] {
				break
			}
			if formNum == nums1 {
				return true
			}
			r++
		}
		l++
	}

	return false
}

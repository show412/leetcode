// https://leetcode.com/problems/longest-palindromic-substring/
/*
Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/
// TC O(n^2) SC O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	lastStart := 0
	lastEnd := 0
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		// for this case like "aba"
		start, end = checkPalindrome(i, i, s)
		if end-start > lastEnd-lastStart {
			lastStart = start
			lastEnd = end
		}
		// for this case like "abba"
		start, end = checkPalindrome(i, i+1, s)
		if end-start > lastEnd-lastStart {
			lastStart = start
			lastEnd = end
		}
	}
	return s[lastStart:(lastEnd + 1)]
}

func checkPalindrome(s int, e int, str string) (int, int) {
	for s >= 0 && e < len(str) && str[s] == str[e] {
		s--
		e++
	}
	return s + 1, e - 1
}

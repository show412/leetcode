// https://leetcode.com/problems/palindrome-permutation/
/*
Given a string, determine if a permutation of the string could form a palindrome.

Example 1:

Input: "code"
Output: false
Example 2:

Input: "aab"
Output: true
Example 3:

Input: "carerac"
Output: true
*/
func canPermutePalindrome(s string) bool {
	m := make(map[byte]bool, 0)
	for i := 0; i < len(s); i++ {
		exist := m[s[i]]
		if exist == true {
			delete(m, s[i])
		} else {
			m[s[i]] = true
		}
	}
	return len(m) <= 1
}

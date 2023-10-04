import "strings"

// https://leetcode.com/problems/valid-palindrome/
/*
Given a string, determine if it is a palindrome,
considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1
	for start != end && start <= end {
		for start < len(s) && checkCase(s[start]) == false {
			start++
		}
		for end >= 0 && checkCase(s[end]) == false {
			end--
		}
		if start <= end && s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func checkCase(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}

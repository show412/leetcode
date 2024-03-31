/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-25 09:54:54
 * @Description: file content
 */
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
/*
key: 1, consider even and odd condition 
2, how to check palindrome, we can start from i to expand its left and right
*/
// TC O(n^2) SC O(1)
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res := ""
	
	for i := 0; i < len(s); i++ {
		// for this case of odd, like "aba"
        start, end := i,i
		start, end = checkPalindrome(i, i, s)
		if end-start+1 > len(res) {
			res = s[start:(end+1)]
		}
		// for this case of even like "abba"
		start, end = checkPalindrome(i, i+1, s)
		if end-start+1 > len(res) {
			res = s[start:(end+1)]
		}
	}
	return res
}

func checkPalindrome(s int, e int, str string) (int, int) {
	for s >= 0 && e < len(str) && str[s] == str[e] {
		s--
		e++
	}
	return s + 1, e - 1
}

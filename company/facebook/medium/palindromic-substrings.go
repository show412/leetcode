// https://leetcode.com/problems/palindromic-substrings/
/*
Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:

Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".


Example 2:

Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".


Note:

The input string length won't exceed 1000.
*/
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(s); i++ {
		// for this case like "aba"
		num1 := checkPalindrome(i, i, s)
		res += num1
		// for this case like "aaa"
		num2 := checkPalindrome(i, i+1, s)
		res += num2
	}
	return res
}
func checkPalindrome(s int, e int, str string) int {
	res := 0
	for s >= 0 && e < len(str) && str[s] == str[e] {
		s--
		e++
		res++
	}
	return res
}

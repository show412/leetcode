// https://leetcode.com/problems/longest-palindromic-subsequence/
/*
Given a string s, find the longest palindromic subsequence's length in s.
You may assume that the maximum length of s is 1000.

Example 1:
Input:

"bbbab"
Output:
4
One possible longest palindromic subsequence is "bbbb".
Example 2:
Input:

"cbbd"
Output:
2
One possible longest palindromic subsequence is "bb".
*/
// 用动态规划解决
func longestPalindromeSubseq(s string) int {

}

// 这种方法不对 要找的是最长回文子串的长度 意思是可以改变子串的顺序
// 比如 bbbab 最长是4 因为 可以bbbb
func longestPalindromeSubseq(s string) int {
	res := 0
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		s1, e1 := checkPalindrome(i, i, s)
		if e1-s1 > end-start {
			start = s1
			end = e1
			res = e1 - s1 + 1
		}
		s1, e1 = checkPalindrome(i, i+1, s)
		if e1-s1 > end-start {
			start = s1
			end = e1
			res = e1 - s1 + 1
		}
	}
	return res
}

func checkPalindrome(s int, e int, str string) (int, int) {
	for s >= 0 && e < len(s) && str[s] == str[e] {
		s--
		e++
	}
	return s + 1, e - 1
}

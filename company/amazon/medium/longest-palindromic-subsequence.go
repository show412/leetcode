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
/*
perhaps there is two cases:
1,aba
2,aa
for aba, if s[i] !=s[j](e.g. a, b), by the max(dp[i+1][j], dp[i][j-1])
it could cover these case
*/
func longestPalindromeSubseq(s string) int {
	// dp[i][j] means from i to j, the max palindrome sub string number
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	// it's from the last of s
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

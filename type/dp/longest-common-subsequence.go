/*
 * @Author: hongwei.sun
 * @Date: 2024-04-07 22:10:00
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-07 22:40:05
 * @Description: file content
 */

 /*
 https://leetcode.com/problems/longest-common-subsequence/description/
 Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

 

Example 1:

Input: text1 = "abcde", text2 = "ace" 
Output: 3  
Explanation: The longest common subsequence is "ace" and its length is 3.
Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
 

Constraints:

1 <= text1.length, text2.length <= 1000
text1 and text2 consist of only lowercase English characters.
 */
/*
这题的关键是定义二维的dp[i][j] 代表 text1的第i个字符和 text2的第j个字符的LCS
最后dp[0][0]就是答案
*/
func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				// 这里是难点，如果这两个字符串不等，那dp[i][j]可以认为 
				// 去掉text1里第i个往后 与 text2里第j个往后的比 或者
				// text1里第i个往后 与 去掉text2里第j个往后的比中的最大
				// 为什么不是dp[i+1][j+1]? 因为这样就少case了
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
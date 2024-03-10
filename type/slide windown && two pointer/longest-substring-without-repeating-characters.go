/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: your name
 * @LastEditTime: 2024-03-10 10:56:56
 * @Description: file content
 */
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
/*
Time complexity : O(2n) = O(n)O(2n)=O(n). In the worst case each character will be visited twice by ii and jj.

Space complexity : O(min(m, n))O(min(m,n)).
Same as the previous approach. We need O(k)O(k) space for the sliding window, where kk is the size of the Set.
The size of the Set is upper bounded by the size of the string nn and the size of the charset/alphabet mm.
*/
func lengthOfLongestSubstring(s string) int {
	res := 0
	l := 0
	r := 0
	m := make(map[byte]int, 0)
	for r < len(s) {
		m[s[r]]++
		for l < r && m[s[r]] > 1 {
			m[s[l]]--
			l++
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// optimize the hashset
func lengthOfLongestSubstring(s string) int {
	res := 0
	l := 0
	r := 0
	m := make(map[byte]int, 0)
	for r < len(s) {
		if _, ok := m[s[r]]; ok {
			l = max(l, m[s[r]]+1)
		}
		m[s[r]] = r
		res = max(res, r-l+1)
		r++
	}
	return res
}
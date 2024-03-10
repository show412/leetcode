/*
 * @Author: hongwei.sun
 * @Date: 2024-03-10 10:42:10
 * @LastEditors: your name
 * @LastEditTime: 2024-03-10 11:46:15
 * @Description: file content
 */
/*
https://leetcode.com/problems/longest-repeating-character-replacement/
You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.



Example 1:

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
Example 2:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.


Constraints:

1 <= s.length <= 105
s consists of only uppercase English letters.
0 <= k <= s.length
*/
/*
sliding window, key is window length minus 出现频率最高的字母的值就是可以替换的次数
如果可以替换的次数小于K 可以向右扩大窗口，否则就缩小窗口
这里的关键是要想到窗口长度减去出现频率最高的就是可以替换的次数
*/
func characterReplacement(s string, k int) int {
	res := 0
	l := 0
	// r := 0
	m := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		m[s[r]]++
		for ((r - l + 1) - maxFrequency(m)) > k {
			m[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxFrequency(m map[byte]int) int {
	maxF := 0
	for _, v := range m {
		maxF = max(maxF, v)
	}
	return maxF
}
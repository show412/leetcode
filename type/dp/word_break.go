/*
 * @Author: hongwei.sun
 * @Date: 2024-04-06 15:04:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-06 16:52:13
 * @Description: file content
 */
//  https://leetcode.com/problems/word-break/description/
/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

 

Example 1:

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
 

Constraints:

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.
*/
/*
要是整个字符串能被切割，意味着从后面开始某个位置能被切割 并且去掉这个单词之前的也能被切割
dp[] 代表在哪里加空格segment是否能work, 长度要算上s的前后 dp[0] = true 我们要看的就是dp[len(s)]是否为true
map记录dict的word
*/
func wordBreak(s string, wordDict []string) bool {
    // 代表在哪里加空格segment是否能work, 要算上s的前后
	/*
		比如 catsandog 这里前后有十个空格算上  字符串里每两个之间有一个，加上前后总共有10个
	*/
	dp := make([]bool, len(s)+1)
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	// initialize dp[0], by default it's true because if there is segmented from 0 index to i index in s
	// means dp[i] = true, we need to make it works 
	dp[0] = true
	// iterate to add space in every single s to check if it can be in one word from 
	for i := 0; i < len(s); i++ {
		for word, _ := range m {
			/*
			如果一个i能被segment 意味着隔着word长度的前面都能被segment
			这里要遍历m 因为可能有多种segment的方式
			*/
			if i >= len(word)-1 && m[s[(i+1-len(word)):(i+1)]] == true && dp[i+1-len(word)] == true {
				dp[i+1] = true
			}
		}
	}
	return dp[len(s)]
}
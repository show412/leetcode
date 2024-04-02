/*
 * @Author: hongwei.sun
 * @Date: 2024-04-02 16:16:47
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 17:24:23
 * @Description: file content
 */
// https://leetcode.com/problems/shortest-palindrome/description/
/*
You are given a string s. You can convert s to a 
palindrome
 by adding characters in front of it.

Return the shortest palindrome you can find by performing this transformation.

 

Example 1:

Input: s = "aacecaaa"
Output: "aaacecaaa"
Example 2:

Input: s = "abcd"
Output: "dcbabcd"
 

Constraints:

0 <= s.length <= 5 * 104
s consists of lowercase English letters only.
*/
/*
直接的想法就是从后开始找到这个字符串中最长的回文的字符串，把后面不是回文的reverse之后接到前面结果就是最短的回文
比如   一个字符串可以分为  a+b
a是回文， b不是回文字符串. 那结果就是bab
*/
func shortestPalindrome(s string) string {
    // 用两个指针指向字符串，初始值是开始和结尾
	l := 0
	r := len(s) - 1
	// to point which position is end of palindrome, initialize is also len(s) -1
	end := len(s) - 1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			// l back to starting to traverse to compare string again
			l = 0
			// end move forward 1 position and assign to r, repeat to compare again
			end--
			r = end
		}
	}
	return reverseString(s[(end+1):]) + s
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
import "strings"

/*
 * @Author: hongwei.sun
 * @Date: 2023-10-09 15:35:19
 * @LastEditors: your name
 * @LastEditTime: 2023-10-09 15:51:34
 * @Description: file content
 https://leetcode.com/problems/length-of-last-word/
 Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring
 consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.


Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.
*/
// this runtime is ok, but memory cost is high
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	s_slice := strings.Split(s, " ")
	return len(s_slice[len(s_slice)-1])
}

// good runtime and good memory cost
func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && res != 0 {
			return res
		} else if s[i] != ' ' {
			res += 1
		}
	}
	return res
}
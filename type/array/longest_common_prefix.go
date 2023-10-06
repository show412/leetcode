import "strings"

/*
 * @Author: hongwei.sun
 * @Date: 2023-10-06 15:29:47
 * @LastEditors: your name
 * @LastEditTime: 2023-10-06 15:47:52
 * @Description: file content
 https://leetcode.com/problems/longest-common-prefix/
 Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/
func longestCommonPrefix(strs []string) string {
	pre := strs[0]
	for i := 0; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], pre) {
			// this is one way to loop strings in go forwards
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}
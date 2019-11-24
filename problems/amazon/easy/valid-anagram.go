// https://leetcode.com/problems/valid-anagram/
/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for j := 0; j < len(t); j++ {
		m[t[j]-'a']--
		if m[t[j]-'a'] < 0 {
			return false
		}
	}
	return true
}

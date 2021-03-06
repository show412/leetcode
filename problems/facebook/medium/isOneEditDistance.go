// https://leetcode.com/problems/one-edit-distance/
/*
Given two strings s and t, determine if they are both one edit distance apart.

Note:

There are 3 possiblities to satisify one edit distance apart:

Insert a character into s to get t
Delete a character from s to get t
Replace a character of s to get t
Example 1:

Input: s = "ab", t = "acb"
Output: true
Explanation: We can insert 'c' into s to get t.
Example 2:

Input: s = "cab", t = "ad"
Output: false
Explanation: We cannot get t from s by only one step.
Example 3:

Input: s = "1203", t = "1213"
Output: true
Explanation: We can replace '0' with '1' to get t.
*/
// notice: if s == t, it should return false
func isOneEditDistance(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls > lt {
		return isOneEditDistance(t, s)
	}
	if lt-ls > 1 {
		return false
	}
	for i := 0; i < ls; i++ {
		if s[i] != t[i] {
			if ls == lt {
				return s[(i+1):] == t[(i+1):]
			} else {
				return s[i:] == t[(i+1):]
			}
		}
	}
	// if s == t, it should return false so it check the ls+1 == lt here
	return ls+1 == lt
}

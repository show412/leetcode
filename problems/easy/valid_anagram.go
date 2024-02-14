import "sort"

/*
 * @Author: hongwei.sun
 * @Date: 2024-02-14 10:29:42
 * @LastEditors: your name
 * @LastEditTime: 2024-02-14 11:22:04
 * @Description: file content

 Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false


Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.


Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/
// can use sort firstly then compare strings
// can use hashmap

func isAnagram(s string, t string) bool {
	s1 := []rune(s)
	t1 := []rune(t)
	sort.SliceStable(s1, func(i int, j int) bool { return s1[i] < s1[j] })
	sort.SliceStable(t1, func(i int, j int) bool { return t1[i] < t1[j] })
	if string(s1) == string(t1) {
		return true
	}
	return false
}
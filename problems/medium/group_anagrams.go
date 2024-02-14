/*
 * @Author: hongwei.sun
 * @Date: 2024-02-14 11:19:11
 * @LastEditors: your name
 * @LastEditTime: 2024-02-14 11:48:19
 * @Description: file content
 */
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

// Constraints:

// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.

// m key is 26 charactor array which value is times of charactor happens
type charCount [26]int

func groupAnagrams(strs []string) [][]string {
	m := make(map[charCount][]string)
	for _, word := range strs {
		var charCount charCount
		for _, w := range word {
			charCount[w-'a']++
		}
		m[charCount] = append(m[charCount], word)
	}
	var ret [][]string
	for _, value := range m {
		ret = append(ret, value)
	}
	return ret
}

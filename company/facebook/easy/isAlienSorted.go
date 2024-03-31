// https://leetcode.com/problems/verifying-an-alien-dictionary/
/*
In an alien language, surprisingly they also use english lowercase letters,
but possibly in a different order. The order of the alphabet is
some permutation of lowercase letters.

Given a sequence of words written in the alien language,
and the order of the alphabet,
return true if and only if the given words are sorted lexicographicaly in this alien language.



Example 1:

Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
Example 2:

Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Output: false
Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
Example 3:

Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
Output: false
Explanation: The first three characters "app" match, and the second string is shorter (in size.)
According to lexicographical rules "apple" > "app", because 'l' > '∅',
where '∅' is defined as the blank character which is less than any other character (More info).


Note:

1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
All characters in words[i] and order are english lowercase letters.
*/
/*
case:
["fxasxpc","dfbdrifhp","nwzgs","cmwqriv","ebulyfyve","miracx","sxckdwzv","dtijzluhts","wwbmnge","qmjwymmyox"]
"zkgwaverfimqxbnctdplsjyohu"
*/
func isAlienSorted(words []string, order string) bool {
	lexi := make(map[byte]int, 0)
	for k, v := range order {
		if _, ok := lexi[byte(v)]; !ok {
			lexi[byte(v)] = k
		}
	}
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]
		start := 0
		for start < len(word1) {
			if start >= len(word2) {
				return false
			}
			if lexi[word1[start]] < lexi[word2[start]] {
				break
			}
			if lexi[word1[start]] > lexi[word2[start]] {
				return false
			}
			start++
		}
	}
	return true
}

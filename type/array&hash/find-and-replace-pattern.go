// https://leetcode.com/problems/find-and-replace-pattern/
/*
You have a list of words and a pattern, and you want to know which words in words matches the pattern.

A word matches the pattern if there exists a permutation of letters p
so that after replacing every letter x in the pattern with p(x), we get the desired word.

(Recall that a permutation of letters is a bijection from letters to letters:
	every letter maps to another letter, and no two letters map to the same letter.)

Return a list of the words in words that match the given pattern.

You may return the answer in any order.



Example 1:

Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
Output: ["mee","aqq"]
Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
since a and b map to the same letter.


Note:

1 <= words.length <= 50
1 <= pattern.length = words[i].length <= 20
*/
// two map solution, if there is no pattern the two maps are contradiction
func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		word := words[i]
		if match(word, pattern) == true {
			res = append(res, word)
		}
	}
	return res
}

func match(word string, pattern string) bool {
	m1 := make(map[byte]byte, 0)
	m2 := make(map[byte]byte, 0)
	for i := 0; i < len(word); i++ {
		if _, ok1 := m1[word[i]]; !ok1 {
			m1[word[i]] = pattern[i]
		}
		if _, ok2 := m2[pattern[i]]; !ok2 {
			m2[pattern[i]] = word[i]
		}
		if m2[pattern[i]] != word[i] || m1[word[i]] != pattern[i] {
			return false
		}
	}
	return true
}

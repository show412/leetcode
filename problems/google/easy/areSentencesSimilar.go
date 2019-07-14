// https://leetcode.com/problems/sentence-similarity/
/*
Given two sentences words1, words2 (each represented as an array of strings), and a list of similar word pairs pairs, determine if two sentences are similar.

For example, "great acting skills" and "fine drama talent" are similar, if the similar word pairs are pairs = [["great", "fine"], ["acting","drama"], ["skills","talent"]].

Note that the similarity relation is not transitive. For example, if "great" and "fine" are similar, and "fine" and "good" are similar, "great" and "good" are not necessarily similar.

However, similarity is symmetric. For example, "great" and "fine" being similar is the same as "fine" and "great" being similar.

Also, a word is always similar with itself. For example, the sentences words1 = ["great"], words2 = ["great"], pairs = [] are similar, even though there are no specified similar word pairs.

Finally, sentences can only be similar if they have the same number of words. So a sentence like words1 = ["great"] can never be similar to words2 = ["doubleplus","good"].

Note:

The length of words1 and words2 will not exceed 1000.
The length of pairs will not exceed 2000.
The length of each pairs[i] will be 2.
The length of each words[i] and pairs[i][j] will be in the range [1, 20].
*/
/*
use case:
["great","acting","skills"]
["fine","drama","talent"]
[["great","fine"],["drama","acting"],["skills","talent"]]
true
*/
func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}
	sMap := make(map[string][]string)
	for i := 0; i < len(pairs); i++ {
		if _, ok := sMap[pairs[i][0]]; ok {
			sMap[pairs[i][0]] = append(sMap[pairs[i][0]], pairs[i][1])
		} else {
			sMap[pairs[i][0]] = []string{pairs[i][1]}
		}
		if _, ok := sMap[pairs[i][1]]; ok {
			sMap[pairs[i][1]] = append(sMap[pairs[i][1]], pairs[i][0])
		} else {
			sMap[pairs[i][1]] = []string{pairs[i][0]}
		}
	}
	start := 0
	for start < len(words1) {
		if words1[start] == words2[start] || existWord(sMap[words1[start]], words2[start]) || existWord(sMap[words2[start]], words1[start]) {
			start++
			continue
		}
		return false
	}
	return true

}

func existWord(strList []string, word string) bool {
	for i := 0; i < len(strList); i++ {
		if strList[i] == word {
			return true
		}
	}
	return false
}

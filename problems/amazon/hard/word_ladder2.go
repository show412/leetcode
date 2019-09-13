// https://leetcode.com/problems/word-ladder-ii/
/*
Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:

Only one letter can be changed at a time
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: []

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

*/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := make(map[string]struct{})
	current, trace := make(map[string]struct{}), make(map[string][]string)
	for _, s := range wordList {
		dict[s] = struct{}{}
		trace[s] = make([]string, 0)
	}
	dict[beginWord] = struct{}{}
	trace[beginWord] = make([]string, 0)
	current[beginWord] = struct{}{}
	_, ok := current[endWord]
	for len(current) != 0 && !ok {
		for word, _ := range current {
			delete(dict, word)
		}
		next := make(map[string]struct{})
		for word, _ := range current {
			for i := range word {
				for _, c := range "abcdefghijklmnopqrstuvwxyz" {
					candidate := word[:i] + string(c) + word[i+1:]
					if _, ok := dict[candidate]; ok {
						trace[candidate] = append(trace[candidate], word)
						next[candidate] = struct{}{}
					}
				}
			}
		}
		current = next
		_, ok = current[endWord]
	}
	var results [][]string
	if len(current) != 0 {
		backtrace(&results, trace, []string{}, endWord)
	}
	return results
}

func backtrace(results *[][]string, trace map[string][]string, path []string, word string) {
	if len(trace[word]) == 0 {
		*results = append(*results, append([]string{word}, path...))
	} else {
		for _, prev := range trace[word] {
			backtrace(results, trace, append([]string{word}, path...), prev)
		}
	}
}

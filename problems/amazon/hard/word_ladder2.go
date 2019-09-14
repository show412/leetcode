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

	result := make([][]string, 0)

	wordMap := make(map[string]bool)

	for _, w := range wordList {
		wordMap[w] = true
	}

	if !wordMap[endWord] {
		return result
	}

	// create a queue, track the path
	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})

	// queueLen is used to track how many slices in queue are in the same level
	// if found a result, I still need to finish checking current level cause I need to return all possible paths
	queueLen := 1
	// use to track strings that this level has visited
	// when queueLen == 0, remove levelMap keys in wordMap
	// levelMap is to prevent the cycle to travers
	levelMap := make(map[string]bool)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		lastWord := path[len(path)-1]

		for i := 0; i < len(lastWord); i++ {
			for c := 'a'; c <= 'z'; c++ {
				nextWord := lastWord[:i] + string(c) + lastWord[i+1:]

				if nextWord == endWord {
					path = append(path, endWord)
					result = append(result, path)
					continue
				}
				if wordMap[nextWord] {
					// different from word ladder, don't remove the word from wordMap immediately
					// same level could reuse the key.
					// delete from wordMap only when currently level is done.
					levelMap[nextWord] = true
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, nextWord)
					queue = append(queue, newPath)

				}
			}
		}

		queueLen--
		// if queueLen is 0, means finish traversing current level. if result is not empty, return result
		if queueLen == 0 {
			if len(result) > 0 {
				return result
			}
			for k, _ := range levelMap {
				delete(wordMap, k)
			}
			// clear levelMap
			levelMap = make(map[string]bool)
			queueLen = len(queue)
		}

	}

	return result
}

import (
	"strings"
)

// Given two words (start and end), and a dictionary, find the shortest transformation sequence from start to end, output the length of the sequence.
// Transformation rule such that:

// Only one letter can be changed at a time
// Each intermediate word must exist in the dictionary. (Start and end words do not need to appear in the dictionary )
// Example
// Example 1:

// Input：start = "a"，end = "c"，dict =["a","b","c"]
// Output：2
// Explanation：
// "a"->"c"
// Example 2:

// Input：start ="hit"，end = "cog"，dict =["hot","dot","dog","lot","log"]
// Output：5
// Explanation：
// "hit"->"hot"->"dot"->"dog"->"cog"
// Notice
// Return 0 if there is no such transformation sequence.
// All words have the same length.
// All words contain only lowercase alphabetic characters.
// https://www.lintcode.com/problem/word-ladder/description
// https://www.jiuzhang.com/solution/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) == 1 {
		return 2
	}
	distance := 1
	length := len(wordList)
	// store those words have been visited
	visit := make(map[string]bool, length)
	// for BFS traverse
	queue := []string{}
	queue = append(queue, beginWord)
	// use hash to store the wordList for better performance
	wordMap := make(map[string]bool, length)
	for z := 0; z < len(wordList); z++ {
		wordMap[wordList[z]] = true
	}
	if wordMap[beginWord] == true {
		visit[beginWord] = true
	}
	// when the endWord is not in the wordList, means there is no answer, return 0
	if wordMap[endWord] != true {
		return 0
	}

	for len(queue) != 0 {
		// fmt.Println("***next bfs***")
		// fmt.Println(queue)

		// key, distance++ should be here
		// because in once change, it maybe transform several words not only one word
		// distance也就是BFS中遍历层数结果的意思 所以应该加到这 遍历queue的前面
		distance++
		// it's the key for the shottest
		// 每次进队的单词 有的单词要经过多次才能变成目标单词 但有的单词可能只需要一次
		// "hot"
		// "dog"
		// ["hot","cog","dog","tot","hog","hop","pot","dot"]
		// hot -> cog -> dog
		// hot -> hog -> cog -> dog
		// 所以第一次bfs 可以变成 tot pot dot
		// 第二次bfs应该从这些里面都试一次看看有没有能达到endWord的 这也是BFS的算法实质
		size := len(queue)
		for k := 0; k < size; k++ {
			word := queue[0]
			queue = queue[1:]
			// word := queue[len(queue)-1]
			// queue = queue[:len(queue)-1]
			// fmt.Println(word)
			// fmt.Println(queue)
			wordArray := strings.Split(word, "")
			for i := 0; i < len(wordArray); i++ {
				tmp := make([]string, len(wordArray))
				copy(tmp, wordArray)
				for j := 'a'; j <= 'z'; j++ {
					tmp[i] = string(j)
					transWord := strings.Join(tmp, "")
					if transWord == endWord {
						return distance
					}
					if visit[transWord] != true && wordMap[transWord] == true {
						queue = append(queue, transWord)
						visit[transWord] = true
					}
				}
			}
		}
	}
	// 当没有遍历到endWord的时候 返回0（即使endWord在wordList里也可能遍历不到）
	return 0
}

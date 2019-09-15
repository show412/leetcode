package main

import (
	"fmt"
	"strings"
)

/*
test case:
[0,0]
-1
true

[0,1,0]
0
false

[23, 2, 6, 4, 7],  k=0
false

[5,0,0]
0
true

{23, 2, 6, 4, 7}, 6
true

[23, 2, 4, 6, 7], 6
true
*/
func main() {
	res := findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(res)
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	path := make([]string, 0)
	result := make([][]string, 0)
	graph := make(map[string][]string, 0)
	dict := make(map[string]bool, 0)
	for _, v := range wordList {
		dict[v] = true
	}
	buildGraph(beginWord, endWord, graph, dict)
	dfs(beginWord, endWord, graph, &path, &result)
	return result
}

func buildGraph(beginWord string, endWord string, graph map[string][]string, dict map[string]bool) {
	visited := make(map[string]bool, 0)
	toVisit := make(map[string]bool, 0)
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	toVisit[beginWord] = true
	foundEnd := false
	for len(queue) != 0 {
		for k, v := range toVisit {
			visited[k] = v
		}
		// visited = toVisit
		toVisit = map[string]bool{}
		count := len(queue)
		for i := 0; i < count; i++ {
			word := queue[0]
			queue = queue[1:]
			children := getNextLevel(word, dict)
			for _, child := range children {
				if child == endWord {
					foundEnd = true
				}
				if _, ok := visited[child]; !ok {
					if _, ok1 := graph[word]; !ok1 {
						graph[word] = make([]string, 0)
					}
					graph[word] = append(graph[word], child)
				}
				// 防止重复寻找
				_, ok2 := visited[child]
				_, ok3 := toVisit[child]
				if !ok2 && !ok3 {
					queue = append(queue, child)
					toVisit[child] = true
				}
			}
		}
		if foundEnd == true {
			break
		}
	}
	return
}

func dfs(curWord string, endWord string, graph map[string][]string, path *[]string, result *[][]string) {
	*path = append(*path, curWord)
	if curWord == endWord {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*result = append(*result, tmp)
	} else if _, ok := graph[curWord]; ok {
		for _, nextWord := range graph[curWord] {
			dfs(nextWord, endWord, graph, path, result)
		}
	}
	*path = (*path)[:len(*path)-1]
	return
}

func getNextLevel(word string, dict map[string]bool) []string {
	result := make([]string, 0)
	wordArray := strings.Split(word, "")
	for i := 0; i < len(wordArray); i++ {
		tmp := make([]string, len(wordArray))
		copy(tmp, wordArray)
		for j := 'a'; j <= 'z'; j++ {
			tmp[i] = string(j)
			transWord := strings.Join(tmp, "")
			if transWord == word {
				continue
			}
			if _, ok := dict[transWord]; ok {
				result = append(result, transWord)
			}
		}
	}
	return result
}

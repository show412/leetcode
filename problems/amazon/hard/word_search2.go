import "fmt"

// https://leetcode.com/problems/word-search-ii/
/*
Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.



Example:

Input:
board = [
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
words = ["oath","pea","eat","rain"]

Output: ["eat","oath"]


Note:

All inputs are consist of lowercase letters a-z.
The values of words are distinct.
*/
// tire solution
// TC is O(n^3)? SC is O(n)
type trienode struct {
	val  rune
	word string
	next map[rune]*trienode
}

func findWords(board [][]byte, words []string) []string {
	// this should be a point, common for struct
	root := &trienode{next: map[rune]*trienode{}}
	for _, w := range words {
		p := root
		for i, b := range w {
			if _, ok := p.next[b]; !ok {
				// notice the defination for next loop
				p.next[b] = &trienode{val: b, next: map[rune]*trienode{}}
			}
			if i == len(w)-1 {
				p.next[b].word = w
			}
			// notice the defination for next loop
			p = p.next[b]
		}
	}

	res := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, &res, root, i, j, map[string]bool{})
		}
	}

	return res
}

func dfs(brd [][]byte, res *[]string, node *trienode, i, j int, visited map[string]bool) {
	// use i and j as the combination key to prevent use letter more than once
	k := fmt.Sprintf("%v_%v", i, j)

	if i < 0 || j < 0 || i >= len(brd) || j >= len(brd[0]) || visited[k] {
		return
	}
	nn, ok := node.next[rune(brd[i][j])]
	if !ok {
		return
	}

	if len(nn.word) > 0 {
		*res = append(*res, nn.word)
		nn.word = ""
	}

	visited[k] = true
	dfs(brd, res, nn, i+1, j, visited)
	dfs(brd, res, nn, i-1, j, visited)
	dfs(brd, res, nn, i, j+1, visited)
	dfs(brd, res, nn, i, j-1, visited)
	visited[k] = false
}

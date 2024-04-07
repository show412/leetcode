/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-07 10:16:59
 * @Description: file content
 */
// https://leetcode.com/problems/design-add-and-search-words-data-structure/
/*

void addWord(word)
bool search(word)
search(word) can search a literal word or a regular expression string
containing only letters a-z or .. A .
means it can represent any one letter.

Example:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
Note:
You may assume that all words are consist of lowercase letters a-z.
*/
/*
test case:
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]

["WordDictionary","addWord","addWord","addWord","addWord","search","search","addWord","search","search","search","search","search","search"]
[[],["at"],["and"],["an"],["add"],["a"],[".at"],["bat"],[".at"],["an."],["a.d."],["b."],["a.d"],["."]]

[null,null,null,null,null,false,false,null,true,true,false,false,true,false]

the complex case:
https://leetcode.com/submissions/detail/254987033/testcase/#

*/
// tire tree solution, and it's a common write code for tire tree
// make(t Type) return same Type. The make built-in function allocates and initializes an object of type slice, map, or chan (only).
type WordDictionary struct {
	// 这里next是个hashmap
    next map[rune]*WordDictionary
	isWord bool
 }
 
 
 func Constructor() WordDictionary {
	// 初始化用make make返回一样的type， new方法返回的是 指向type的指针
	 return WordDictionary{next: make(map[rune]*WordDictionary), isWord: false}
 }
 
 
 func (this *WordDictionary) AddWord(word string)  {
	 for _, w := range word {
		if this.next[w] == nil {
			// 初始化用make make返回一样的type， new方法返回的是 指向type的指针
			this.next[w] = &WordDictionary{next: make(map[rune]*WordDictionary), isWord: false}
		}
		this = this.next[w]
	 }
	 this.isWord = true
 }
 
 
 func (this *WordDictionary) Search(word string) bool {
	for k, w := range word {
		if w == '.' {
            // 遍历当前level的所有next里的字母
			for _, d := range this.next {
				// 这是这道题的难点，得用递归去继续找‘.’这种情况
				// 如果能找到是一个词(isWord == true) 就返回true
				// 这里不能用return d.Search, 因为上面是个遍历，这里return的话就直接return第一个了
				// 都遍历完了没有任何一个isWord再返回false
				if d.Search(word[(k+1):]) {
                    return true
                }
			}
            return false
		} else if this.next[w] == nil {
			return false
		} else {
            this = this.next[w]
        }
	}
	return this.isWord
 }
import (
	"regexp"
	"strings"
)

// https://leetcode.com/problems/add-and-search-word-data-structure-design/
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
// we use the regex, it maybe trigger the LTE for the complex case
// there is a tire solution for this more fase
// https://leetcode.com/problems/add-and-search-word-data-structure-design/discuss/341154/go%3A-simple-solution
type WordDictionary struct {
	dict map[string]bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{dict: make(map[string]bool, 0)}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.dict[word] = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	strArray := make([]string, 0)
	flag := false
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {

			strArray = append(strArray, "(.)")

			flag = true
		} else {
			strArray = append(strArray, string(word[i]))
		}
	}
	str := strings.Join(strArray, "")
	// fmt.Println(str)
	if flag == false {
		return this.dict[str]
	} else {
		str = "^" + str + "$"
	}
	for k, _ := range this.dict {
		match, _ := regexp.MatchString(str, k)
		if match == true {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
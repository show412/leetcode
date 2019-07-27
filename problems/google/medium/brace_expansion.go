import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/brace-expansion/
/*
A string S represents a list of words.

Each letter in the word has 1 or more options.  If there is one option, the letter is represented as is.  If there is more than one option, then curly braces delimit the options.  For example, "{a,b,c}" represents options ["a", "b", "c"].

For example, "{a,b,c}d{e,f}" represents the list ["ade", "adf", "bde", "bdf", "cde", "cdf"].

Return all words that can be formed in this manner, in lexicographical order.



Example 1:

Input: "{a,b}c{d,e}f"
Output: ["acdf","acef","bcdf","bcef"]
Example 2:

Input: "abcd"
Output: ["abcd"]


Note:

1 <= S.length <= 50
There are no nested curly brackets.
All characters inside a pair of consecutive opening and ending curly brackets are different.
*/
// 这个不能叫dfs 应该叫递归吧
func expand(S string) []string {
	res := make([]string, 0)
	expandStr(S, &res)
	return res
}

func expandStr(S string, res *[]string) {
	if strings.Index(S, "{") == -1 {
		*res = append(*res, S)
		return
	}
	openIndex := strings.Index(S, "{")
	closeIndex := strings.Index(S, "}")
	// string[:]切分后仍然是string
	// strings.Split 之后仍然是[]string
	options := strings.Split(S[openIndex+1:closeIndex], ",")
	sort.Strings(options)
	for i := 0; i < len(options); i++ {
		expandStr(S[:openIndex]+options[i]+S[closeIndex+1:], res)
	}
}

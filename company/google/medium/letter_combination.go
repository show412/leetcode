import "strings"

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
/*
Given a string containing digits from 2-9 inclusive,
return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons)
is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"}}
	res := make([]string, 0)
	combination := make([]string, 0)
	dfs(m, digits, combination, &res)
	return res
}

func dfs(m map[string][]string, disgits string, combination []string, res *[]string) {
	if disgits == "" || len(disgits) == 0 {
		*res = append(*res, strings.Join(combination, ""))
		return
	}
	number := string(disgits[0])
	disgits = disgits[1:]
	set := m[number]
	for i := 0; i < len(set); i++ {
		combination = append(combination, set[i])
		dfs(m, disgits, combination, res)
		combination = combination[:len(combination)-1]
	}
}

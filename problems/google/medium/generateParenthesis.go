// https://leetcode.com/problems/generate-parentheses/
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// For example, given n = 3, a solution set is:

// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]
// DFS like permutation
// "(" means 1 ")" means -1
// if sum >=0 it's valid
func generateParenthesis(n int) []string {

	dfs(nums, visit, &entry, &res)
}

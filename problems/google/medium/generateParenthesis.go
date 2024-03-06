/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-03-06 11:10:26
 * @Description: file content
 */
import "strings"

// https://leetcode.com/problems/generate-parentheses/
// Given n pairs of parentheses,
// write a function to generate all combinations of well-formed parentheses.

// For example, given n = 3, a solution set is:

// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

/* solution of answer
three conditions to be reseanable
1, when open bracket size less than n, add (
2, when only the open bracket size is bigger than the close bracket size, can )
3, when open bracket and clost bracket are both 3, one result generated
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := make([]string, 0)
	str := ""
	recursionAdd(n, str, 0, 0, &res)
	return res
}
func recursionAdd(n int, str string, open int, close int, res *[]string) {
	if len(str) == 2*n {
		*res = append(*res, str)
		return
	}
	// it seems a DFS recursion
	// here str is almost one stack.
	if open < n {
		recursionAdd(n, str+"(", open+1, close, res)
	}
	if close < open {
		recursionAdd(n, str+")", open, close+1, res)
	}
	return
}

// DFS actually it's almost same solution as above, just we use one stack

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	stack := make([]string, 0)
	result := make([]string, 0)

	dfs(0, 0, n, &stack, &result)

	return result
}

func dfs(openCount int, closeCount int, n int, stack *[]string, result *[]string) {
	if openCount == n && closeCount == n {
		*result = append(*result, strings.Join(*stack, ""))
	}
	if openCount < n {
		*stack = append(*stack, "(")
		dfs(openCount+1, closeCount, n, stack, result)
		*stack = (*stack)[:len(*stack)-1]
	}
	if openCount > closeCount {
		*stack = append(*stack, ")")
		dfs(openCount, closeCount+1, n, stack, result)
		*stack = (*stack)[:len(*stack)-1]
	}
}

// the solution is like divided conquer
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
		return res
	}
	for c := 0; c < n; c++ {
		for _, left := range generateParenthesis(c) {
			for _, right := range generateParenthesis(n - 1 - c) {
				res = append(res, "("+left+")"+right)
			}
		}
	}
	return res
}

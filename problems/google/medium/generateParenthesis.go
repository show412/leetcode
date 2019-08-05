import "strconv"

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

// solution of answer
// because only the open bracket size is bigger than the close bracket size
// it will be reseanable

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
	if open < n {
		recursionAdd(n, str+"(", open+1, close, res)
	}
	if close < open {
		recursionAdd(n, str+")", open, close+1, res)
	}
	return
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
			for _, right := range generateParenthesis(n-1-c) {
				res = append(res, "(" + left + ")" + right)
			}
		}
	}
	return res
}

// the solution is like divided conquer
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> ans = new ArrayList();
        if (n == 0) {
            ans.add("");
        } else {
            for (int c = 0; c < n; ++c)
                for (String left: generateParenthesis(c))
                    for (String right: generateParenthesis(n-1-c))
                        ans.add("(" + left + ")" + right);
        }
        return ans;
    }
}

// DFS like permutation
// "(" means 1 ")" means -1
// if sum >=0 it's valid
// It's TLE
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	nums := make([]int, 0)
	visit := make(map[int]bool)
	uniq := make(map[string]bool)
	res := make([][]int, 0)
	entry := make([]int, 0)
	result := make([]string, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, 1)
		nums = append(nums, -1)
	}
	dfs(nums, visit, entry, &res, uniq)
	for _, item := range res {
		str := ""
		for _, i := range item {
			if i == 1 {
				str += "("
			} else if i == -1 {
				str += ")"
			}
		}
		result = append(result, str)
	}
	return result
}

func dfs(nums []int, visit map[int]bool, entry []int, res *[][]int, uniq map[string]bool) {
	if len(entry) == len(nums) {
		cpy := make([]int, len(entry))
		copy(cpy, entry)
		str := ""
		for i := 0; i < len(cpy); i++ {
			str += strconv.Itoa(cpy[i])
		}
		// fmt.Println(str)
		if v, ok := uniq[str]; ok && v == true {
			return
		}
		*res = append(*res, cpy)
		uniq[str] = true
		return
	}
	for i := 0; i < len(nums); i++ {
		// if i != 0 && nums[i-1] == nums[i] {
		// 	continue
		// }
		if v, ok := visit[i]; ok && v == true {
			continue
		}
		entry = append(entry, nums[i])
		visit[i] = true
		if isValid(entry) == true {
			dfs(nums, visit, entry, res, uniq)
		}
		entry = entry[:len(entry)-1]
		visit[i] = false
	}
}

func isValid(arr []int) bool {
	cur := 0
	for _, i := range arr {
		cur += i
		if cur < 0 {
			return false
		}
	}
	return true
}

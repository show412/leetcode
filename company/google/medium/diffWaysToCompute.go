import "strconv"

// https://leetcode.com/problems/different-ways-to-add-parentheses/
/*
Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. The valid operators are +, - and *.

Example 1:

Input: "2-1-1"
Output: [0, 2]
Explanation:
((2-1)-1) = 0
(2-(1-1)) = 2
Example 2:

Input: "2*3-4*5"
Output: [-34, -14, -10, -10, 10]
Explanation:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
*/
/*
It's a divided-conquer, all possibile is divided int two parts
then conquer each result
*/
func diffWaysToCompute(input string) []int {
	res := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == '-' || input[i] == '+' || input[i] == '*' {
			part1 := ([]byte(input))[:i]
			part2 := ([]byte(input))[(i + 1):]
			part1Res := diffWaysToCompute(string(part1))
			part2Res := diffWaysToCompute(string(part2))
			for _, p1 := range part1Res {
				for _, p2 := range part2Res {
					c := 0
					if input[i] == '+' {
						c = p1 + p2
					} else if input[i] == '-' {
						c = p1 - p2
					} else if input[i] == '*' {
						c = p1 * p2
					}
					res = append(res, c)
				}
			}
		}
	}
	// notice this should be len(res) == 0,
	// it is to make the recusion to the last step for one single number
	if len(res) == 0 {
		v, _ := strconv.Atoi(input)
		res = append(res, v)
	}
	return res
}


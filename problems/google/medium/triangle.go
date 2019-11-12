import "math"

// https://leetcode.com/problems/triangle/
/* Given a triangle,
find the minimum path sum from top to bottom.
Each step you may move to adjacent numbers on the row below.

Example
Example 1:

Input the following triangle:
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
Output: 11
Explanation: The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
Example 2:

Input the following triangle:
[
     [2],
    [3,2],
   [6,5,7],
  [4,4,8,1]
]
Output: 12
Explanation: The minimum path sum from top to bottom is 12 (i.e., 2 + 2 + 7 + 1 = 12).
Notice
Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
*/
/**
 * @param triangle: a list of lists of integers
 * @return: An integer, minimum path sum
 */
func minimumTotal(triangle [][]int) int {
	// write your code here
	// DP to solve the issue
	f := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		f[i] = make([]int, len(triangle[i]))
	}
	res := int(^uint(0) >> 1)
	f[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		f[i][0] += f[i-1][0] + triangle[i][0]
	}

	for i := 1; i < len(triangle); i++ {
		for j := 1; j < len(triangle[i]); j++ {
			if j == len(triangle[i])-1 {
				f[i][j] = triangle[i][j] + f[i-1][j-1]
			} else {
				f[i][j] = triangle[i][j] + int(math.Min(float64(f[i-1][j-1]), float64(f[i-1][j])))
			}

		}
	}
	bottom := len(triangle) - 1
	// fmt.Println(triangle[bottom])
	// fmt.Println(f)
	for j := 0; j < len(triangle[bottom]); j++ {
		if res > f[bottom][j] {
			res = f[bottom][j]
		}
	}
	return res
}

// https://www.lintcode.com/problem/binary-tree-maximum-path-sum/description
// Given a binary tree, find the maximum path sum.

// The path may start and end at any node in the tree.

// Example
// Example 1:
// 	Input:  For the following binary tree（only one node）:
// 	2
// 	Output：2

// Example 2:
// 	Input:  For the following binary tree:

//       1
//      / \
//     2   3

// 	Output: 6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: The root of binary tree.
 * @return: An integer
 */
//  singlePath means the max path to be through root
// maxPath means the max path unnessary to be through root
type ResultType struct {
	singlePath int
	maxPath    int
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{singlePath: INT_MIN, maxPath: INT_MIN}
	}
	left := helper(root.Left)
	right := helper(root.Right)
	singlePath := max(0, max(left.singlePath, right.singlePath)) + root.Val
	maxPath := max(left.maxPath, right.maxPath)
	maxPath = max(maxPath, max(left.singlePath, 0)+max(right.singlePath, 0)+root.Val)
	return ResultType{singlePath: singlePath, maxPath: maxPath}
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxPathSum(root *TreeNode) int {
	// write your code here
	result := helper(root)
	return result.maxPath
}

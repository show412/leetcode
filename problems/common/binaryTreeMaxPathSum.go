import "math"

// https://leetcode.com/problems/binary-tree-maximum-path-sum/
/*
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node
to any node in the tree along the parent-child connections.
The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6
Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// singlePath means the max path to be through root
// maxPath means the max path unnessary to be through root,but at least one node here
// this struct is KEY
// 思路就是 最大路径要不就通过 root 要不就在左子树或者右子树
type ResultType struct {
	singlePath int
	maxPath    int
}

const INT_MIN = math.MinInt64

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{singlePath: INT_MIN, maxPath: INT_MIN}
	}
	// Divide
	left := helper(root.Left)
	right := helper(root.Right)
	// Conquer
	singlePath := max(0, max(left.singlePath, right.singlePath)) + root.Val
	maxPath := max(left.maxPath, right.maxPath)
	// maxPath is in the maxPath of left or the maxPath of right or through the root
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

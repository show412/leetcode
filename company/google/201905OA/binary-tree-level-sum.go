// Given a binary tree and an integer which is the depth of the target level.

// Calculate the sum of the nodes in the target level.

// Example
// Example 1:

// Input：{1,2,3,4,5,6,7,#,#,8,#,#,#,#,9},2
// Output：5
// Explanation：
//      1
//    /   \
//   2     3
//  / \   / \
// 4   5 6   7
//    /       \
//   8         9
// 2+3=5
// Example 2:

// Input：{1,2,3,4,5,6,7,#,#,8,#,#,#,#,9},3
// Output：22
// Explanation：
//      1
//    /   \
//   2     3
//  / \   / \
// 4   5 6   7
//    /       \
//   8         9
// 4+5+6+7=22

// https://www.lintcode.com/problem/binary-tree-level-sum/description
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: the root of the binary tree
 * @param level: the depth of the target level
 * @return: An integer
 */
func levelSum(root *TreeNode, level int) int {
	// write your code here
	if root == nil {
		return 0
	}
	if level == 1 {
		return root.Val
	}
	leftResult := 0
	rightResult := 0
	if root.Left != nil {
		leftResult = levelSum(root.Left, level-1)
	}
	if root.Right != nil {
		rightResult = levelSum(root.Right, level-1)
	}
	result := leftResult + rightResult
	return result
}

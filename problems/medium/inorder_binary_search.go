// https://leetcode.com/problems/binary-tree-inorder-traversal/
/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  there are two other solution without recursion
// https://leetcode.com/problems/binary-tree-inorder-traversal/solution/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	if len(left) != 0 {
		res = append(res, left...)
	}
	res = append(res, root.Val)
	if len(right) != 0 {
		res = append(res, right...)
	}
	return res
}

// with stack, TC O(n), SC O(n)
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

// there is another solution with Morris Traversal in https://leetcode.com/articles/binary-tree-inorder-traversal/
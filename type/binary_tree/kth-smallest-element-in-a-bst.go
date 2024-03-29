// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
/*
Given a binary search tree, write a function kthSmallest
to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:

Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
Follow up:
What if the BST is modified (insert/delete operations)
often and you need to find the kth smallest frequently?
How would you optimize the kthSmallest routine?
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  https://leetcode.com/problems/kth-smallest-element-in-a-bst/solution/
// TC O(H+k) SC O(H+k)
// 一般BST找第k个值都可以这么处理 减小时间复杂度
func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return root.Val
}

func kthSmallest(root *TreeNode, k int) int {
	res := inorderTraversal(root)
	return res[k-1]
}

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

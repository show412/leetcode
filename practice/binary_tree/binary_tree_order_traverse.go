// https://leetcode.com/problems/binary-tree-level-order-traversal/
/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  it's a iteration BFS solution
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	level := make([]*TreeNode, 0)
	level = append(level, root)
	res := make([][]int, 0)
	for len(level) > 0 {
		levelVal := make([]int, 0)
		nextLevel := make([]*TreeNode, 0)
		for i := 0; i < len(level); i++ {
			levelVal = append(levelVal, level[i].Val)
		}
		res = append(res, levelVal)
		for i := 0; i < len(level); i++ {
			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}
		level = nextLevel
	}
	return res
}

// it's a recursion solution

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	helper(root, 0, &res)
	return res
}

func helper(node *TreeNode, level int, res *[][]int) {
	if len(*res) == level {
		(*res) = append(*res, make([]int, 0))
	}
	(*res)[level] = append((*res)[level], node.Val)
	if node.Left != nil {
		helper(node.Left, level+1, res)
	}
	if node.Right != nil {
		helper(node.Right, level+1, res)
	}
	return
}

// https://leetcode.com/problems/binary-tree-right-side-view/
/*
It's a template for bfs refer to
Given a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Example:

Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
https://leetcode.com/problems/binary-tree-right-side-view/
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		value := 0
		for i := 0; i < size; i++ {
			node := queue[0]
			if size > 1 {
				queue = queue[1:]
			} else {
				queue = []*TreeNode{}
			}
			value = node.Val
			// it's the key for this problem
			if i == size-1 {
				res = append(res, value)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

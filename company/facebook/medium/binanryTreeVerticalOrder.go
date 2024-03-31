// https://leetcode.com/problems/binary-tree-vertical-order-traversal/
/*
Given a binary tree, return the vertical order traversal of its nodes' values.
(ie, from top to bottom, column by column).

If two nodes are in the same row and column, the order should be from left to right.

Examples 1:

Input: [3,9,20,null,null,15,7]

   3
  /\
 /  \
 9  20
    /\
   /  \
  15   7

Output:

[
  [9],
  [3,15],
  [20],
  [7]
]
Examples 2:

Input: [3,9,8,4,0,1,7]

     3
    /\
   /  \
   9   8
  /\  /\
 /  \/  \
 4  01   7

Output:

[
  [4],
  [9],
  [3,0,1],
  [8],
  [7]
]
Examples 3:

Input: [3,9,8,4,0,1,7,null,null,null,2,5]
(0's right child is 2 and 1's left child is 5)

     3
    /\
   /  \
   9   8
  /\  /\
 /  \/  \
 4  01   7
    /\
   /  \
   5   2

Output:

[
  [4],
  [9,5],
  [3,0,1],
  [8,2],
  [7]
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
//  用 min 记录偏移量也是一个 common 的做法
func verticalOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	m := make(map[int][]int, 0)
	queue := make([]*TreeNode, 0)
	weight := make(map[*TreeNode]int, 0)
	queue = append(queue, root)
	weight[root] = 0
	min := 0
	for len(queue) != 0 {
		// 这个很关键是把头的拿出来 因为后面的答案要保证顺序
		node := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []*TreeNode{}
		}

		if w, ok := weight[node]; ok {
			if _, ok1 := m[w]; !ok1 {
				m[w] = make([]int, 0)
			}
			m[w] = append(m[w], node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
				weight[node.Left] = w - 1
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				weight[node.Right] = w + 1
			}
			min = MathMin(min, w)
		}
	}

	for len(m[min]) != 0 {
		res = append(res, m[min])
		min++
	}
	return res
}
func MathMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
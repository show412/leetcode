// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if K == 0 {
		return []int{target.Val}
	}
	parent := make(map[*TreeNode]*TreeNode, 0)
	visited := make(map[*TreeNode]bool, 0)
	getParent(root, nil, parent)
	stack := make([]*TreeNode, 0)
	// 先放target 后放nil level初始化就应该是1 因为达到K的时候 当前的符合的level还应该在stack里
	// 用nil进行分层
	stack = append(stack, target)
	stack = append(stack, nil)
	level := 1
	visited[target] = true
	visited[nil] = true
	res := make([]int, 0)
	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			if level == K {
				for i := 0; i < len(stack); i++ {
					if stack[i] != nil {
						res = append(res, stack[i].Val)
					}
				}
				return res
			}
			stack = append(stack, nil)
			level++
		} else {
			if _, ok := visited[node.Left]; !ok && node.Left != nil {
				visited[node.Left] = true
				stack = append(stack, node.Left)
			}
			if _, ok := visited[node.Right]; !ok && node.Right != nil {
				visited[node.Right] = true
				stack = append(stack, node.Right)
			}
			pNode := parent[node]
			if _, ok := visited[pNode]; !ok {
				visited[pNode] = true
				stack = append(stack, pNode)
			}
		}
	}
	return res
}

func getParent(node *TreeNode, pNode *TreeNode, parent map[*TreeNode]*TreeNode) {
	parent[node] = pNode
	if node.Left != nil {
		getParent(node.Left, node, parent)
	}
	if node.Right != nil {
		getParent(node.Right, node, parent)
	}
	return
}

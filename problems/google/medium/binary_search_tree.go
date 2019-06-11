https://leetcode.com/problems/binary-search-tree-iterator
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	 queue []int
}


func Constructor(root *TreeNode) BSTIterator {
	queue := inOrder(root)
	return BSTIterator{queue:queue}
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]int, 0)

	// divided
	left := inOrder(root.Left)
	right := inOrder(root.Right)

	// conquer
	queue = append(queue, left...)
	queue = append(queue, root.Val)
	queue = append(queue, right...)
	return queue
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	queue := 0
	if len(this.queue) > 0 {
		queue = this.queue[0]
		this.queue = this.queue[1:]
	}
	return queue
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.queue) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
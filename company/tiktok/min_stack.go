/*
 * @Author: hongwei.sun
 * @Date: 2024-04-17 20:37:00
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-17 20:37:00
 * @Description: file content
 */
// https://leetcode.com/problems/min-stack/
/*
Design a stack that supports push, pop, top,
and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/

/** initialize your data structure here. */
/*
这个问题的关键是维护一个minimum array去存当时stack里的最小元素，即使有重复的。stack pop的时候同时minmum也pop出来
*/
type MinStack struct {
    stack  []int
	minimum []int
}

func Constructor() MinStack {
	return MinStack{
		stack:   []int{},
		minimum: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minimum) == 0 {
		this.minimum = append(this.minimum, x)
	} else {
		this.minimum = append(this.minimum, min(this.minimum[len(this.minimum)-1], x))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minimum = this.minimum[:len(this.minimum)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minimum[len(this.minimum)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 维持一个栈 维持一个堆就可以了
type MinStack struct {
	stack []int
	heap  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}, heap: []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	cpy := make([]int, len(this.stack))
	copy(cpy, this.stack)
	this.heap = cpy
	HeapSort(this.heap)
}

func HeapSort(c []int) {
	var n = len(c) - 1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, c)
	}
	for end := n; end >= 0; end-- {
		if c[0] < c[end] {
			c[0], c[end] = c[end], c[0]
			minHeap(0, end-1, c)
		}
	}
}

func minHeap(root int, end int, c []int) {
	for {
		var child = 2*root + 1
		//判断是否存在child节点
		if child > end {
			break
		}
		//判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	cpy := make([]int, len(this.stack))
	copy(cpy, this.stack)
	this.heap = cpy
	HeapSort(this.heap)
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	min := 0
	if len(this.heap) != 0 {
		min = this.heap[len(this.heap)-1]
	}

	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

//  better TC and SC with two stack
type MinStack struct {
	stack   []int
	minimum []int
}

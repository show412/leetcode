// https://leetcode.com/problems/max-stack/
/*
Design a max stack that supports push, pop, top, peekMax and popMax.

push(x) -- Push element x onto stack.
pop() -- Remove the element on top of the stack and return it.
top() -- Get the element on the top.
peekMax() -- Retrieve the maximum element in the stack.
popMax() -- Retrieve the maximum element in the stack, and remove it. If you find more than one maximum elements, only remove the top-most one.
Example 1:
MaxStack stack = new MaxStack();
stack.push(5);
stack.push(1);
stack.push(5);
stack.top(); -> 5
stack.popMax(); -> 5
stack.top(); -> 1
stack.peekMax(); -> 5
stack.pop(); -> 1
stack.top(); -> 5
*/

/*
case:
["MaxStack","push","peekMax","popMax"]
[[],[5],[],[]]



*/

type Node struct {
	value int
	index int
}

type MaxStack struct {
	stack    []int
	maxStack []*Node
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{maxStack: nil, stack: nil}
}

// 各种函数名和变量名的写法要学学
func buildMaxStack(stack []int) []*Node {
	if len(stack) == 0 {
		return nil
	}
	max := stack[0]
	maxIndex := 0
	var res []*Node
	for i := 0; i < len(stack); i++ {
		node := &Node{value: max, index: maxIndex}
		if stack[i] >= max {
			node.value = stack[i]
			node.index = i
			max = stack[i]
			maxIndex = i
		}
		res = append(res, node)
	}
	return res
}

func (this *MaxStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.maxStack = buildMaxStack(this.stack)
}

func (this *MaxStack) Pop() int {
	pop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.maxStack = buildMaxStack(this.stack)
	return pop
}

func (this *MaxStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MaxStack) PeekMax() int {
	return this.maxStack[len(this.maxStack)-1].value
}

func (this *MaxStack) PopMax() int {
	max := this.maxStack[len(this.maxStack)-1].value
	index := this.maxStack[len(this.maxStack)-1].index
	this.stack = append(this.stack[:index], this.stack[(index+1):]...)
	this.maxStack = buildMaxStack(this.stack)
	return max
}

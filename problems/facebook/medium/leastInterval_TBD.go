import "container/heap"

// https://leetcode.com/problems/task-scheduler/
/*
Given a char array representing tasks CPU need to do.
It contains capital letters A to Z where different letters represent different tasks.
Tasks could be done without original order. Each task could be done in one interval.
For each interval, CPU could finish one task or just be idle.

However, there is a non-negative cooling interval n that means between two same tasks,
there must be at least n intervals that CPU are doing different tasks or just be idle.

You need to return the least number of intervals the CPU will take to finish
all the given tasks.

Example:

Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation: A -> B -> idle -> A -> B -> idle -> A -> B.

Note:

The number of tasks is in the range [1, 10000].
The integer n is in the range [0, 100].
*/
type Node struct {
	nextCycle, count int
}

type Nodes []*Node

func (n *Nodes) Push(v interface{}) {
	*n = append(*n, v.(*Node))
}

func (n *Nodes) Pop() interface{} {
	rtn := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return rtn
}

func (n *Nodes) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i].nextCycle < n[j].nextCycle || n[i].nextCycle == n[j].nextCycle && n[i].count > n[j].count
}

func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]*Node)

	for _, t := range tasks {
		if m[t] == nil {
			m[t] = &Node{0, 1}
		} else {
			m[t].count++
		}
	}

	nodes := Nodes{}
	for _, task := range m {
		nodes = append(nodes, task)
	}

	heap.Init(&nodes)

	cycle := 0

	for len(nodes) > 0 {
		if nodes[0].nextCycle <= cycle {
			nodes[0].count--
			if nodes[0].count == 0 {
				heap.Pop(&nodes)
			} else {
				nodes[0].nextCycle += n + 1
				heap.Fix(&nodes, 0)
			}
		}
		cycle++
	}

	return cycle
}

/*
Algorithm

Instead of making use of sorting as done in the last approach,
we can also make use of a Max-Heap(queuequeue) to pick the order
in which the tasks need to be executed.
But we need to ensure that the heapification occurs only
after the intervals of cooling time, nn, as done in the last approach.

To do so, firstly, we put only those elements from mapmap into the queuequeue
which have non-zero number of instances.
Then, we start picking up the largest task from the queuequeue for current execution.
(Again, at every instant, we update the current timetime as well.)
We pop this element from the queuequeue.
We also decrement its pending number of instances and
if any more instances of the current task are pending,
we store them(count) in a temporary temptemp list,
to be added later on back into the queuequeue.
We keep on doing so, till a cycle of cooling time has been finished.
After every such cycle,
we add the generated temptemp list back to the queuequeue
for considering the most critical task again.

We keep on doing so till the queuequeue(and temptemp) become totally empty.
At this instant, the current value of timetime gives the required result.


*/

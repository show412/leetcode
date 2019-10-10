import (
	"container/heap"
	"sort"
)

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
/*
AAAABBBEEFFGG 3

A出现了4次，最多，mx=4，那么可以分为mx-1=3块，如下：

A---A---A---

每块有n+1=4个，最后还要加上末尾的一个A，也就是25-24=1个任务，最终结果为13：

ABEFABEGABFGA

再来看另一个例子：

ACCCEEE 2

C和E都出现了3次，最多，mx=3，那么可以分为mx-1=2块，如下：

CE-CE-

每块有n+1=3个，最后还要加上末尾的一个CE，也就是25-23=2个任务，最终结果为8：

CEACE-CE

好，那么此时你可能会有疑问，为啥还要跟原任务个数len相比，取较大值呢？我们再来看一个例子：

AAABBB 0

A和B都出现了3次，最多，mx=3，那么可以分为mx-1=2块，如下：

ABAB

每块有n+1=1个？你会发现有问题，这里明明每块有两个啊，为啥这里算出来n+1=1呢，
因为给的n=0，这有没有矛盾呢，没有！
因为n表示相同的任务间需要间隔的个数，那么既然这里为0了，说明相同的任务可以放在一起，
这里就没有任何限制了，我们只需要执行完所有的任务就可以了，
所以我们最终的返回结果一定不能小于任务的总个数len的，这就是要对比取较大值的原因了。
*/
// refer to https://www.cnblogs.com/grandyang/p/7098764.html
func leastInterval(tasks []byte, n int) int {
	cnt := make([]int, 26)
	for _, task := range tasks {
		cnt[task-'A']++
	}
	sort.Ints(cnt)
	i, maxCount := 25, cnt[25]
	// 拿出现最多byte的个数
	for i >= 0 && cnt[i] == maxCount {
		i--
	}
	// 25-i的意思是最后要加上多少个byte 加上的应该是再加一次出现最多的byte个数
	return max(len(tasks), (maxCount-1)*(n+1)+25-i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 这个解法正确 但是不是很好理解 唯一学到的就是golang heap的使用 至于为什么最小堆可以解出来还得想想
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

https://www.cnblogs.com/grandyang/p/7098764.html

*/

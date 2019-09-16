import "container/heap"

// https://leetcode.com/problems/top-k-frequent-elements/
/*
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n),
where n is the array's size.
*/

type item struct {
	value int
	count int
}
type ItemHeap []item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// heap 堆也要实现push和pop的方法 用指针
func (h *ItemHeap) Push(x interface{}) {
	// 断言 断言只是在接口和类型之间的转换 这里参数是interface
	*h = append(*h, x.(item))
}

// 注意拿最小堆的最小值是h[0], pop出来却是最后一个
// 因为在调用的时候 会先把第0和n-1个交换再拿出n-1个 所以这里只要实现把最后一个pop出来
// refer to https://studygolang.com/articles/13173#43-heappop
func (h *ItemHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func topKFrequent(nums []int, k int) []int {
	heapItem := &ItemHeap{}
	hashItem := make(map[int]int, 0)
	res := make([]int, 0)
	// init heap
	heap.Init(heapItem)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := hashItem[num]; ok {
			hashItem[num] += 1
		} else {
			hashItem[num] = 1
		}
	}

	for key, value := range hashItem {
		h := item{key, -value}
		heap.Push(heapItem, item(h))
	}

	count := 0
	for count < k {
		// 断言 断言只是在接口和类型之间的转换 这里返回的是interface
		res = append(res, heap.Pop(heapItem).(item).value)
		if len(*heapItem) == 0 {
			break
		}
		count++
	}
	return res
}

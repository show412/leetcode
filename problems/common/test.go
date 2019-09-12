package main

import (
	"container/heap"
	"fmt"
)

/*
test case:
[0,0]
-1
true

[0,1,0]
0
false

[23, 2, 6, 4, 7],  k=0
false

[5,0,0]
0
true

{23, 2, 6, 4, 7}, 6
true

[23, 2, 4, 6, 7], 6
true
*/
func main() {
	res := topKFrequent([]int{1}, 1)
	fmt.Println(res)
}

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
		res = append(res, heap.Pop(heapItem).(item).value)
		if len(*heapItem) == 0 {
			break
		}
		count++
	}
	return res
}

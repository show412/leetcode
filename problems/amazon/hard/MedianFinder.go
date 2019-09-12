import "container/heap"

// https://leetcode.com/problems/find-median-from-data-stream/
/*
Median is the middle value in an ordered integer list.
If the size of the list is even, there is no middle value.
So the median is the mean of the two middle value.

For example,
[2,3,4], the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:

void addNum(int num) - Add a integer number from the data stream to the data structure.
double findMedian() - Return the median of all elements so far.


Example:

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2


Follow up:

If all integer numbers from the stream are between 0 and 100,
how would you optimize it?
If 99% of all integer numbers from the stream are between 0 and 100,
how would you optimize it?
*/
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// 这是一个最小堆的实现方法
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// heap 堆也要实现push和pop的方法 用指针
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 注意拿最小堆的最小值是h[0], pop出来却是最后一个
// 因为在调用的时候 会先把第0和n-1个交换再拿出n-1个 所以这里只要实现把最后一个pop出来
// refer to https://studygolang.com/articles/13173#43-heappop
func (h *IntHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

type MedianFinder struct {
	maxHeap, minHeap *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	// h := &IntHeap{}
	max := &IntHeap{}
	min := &IntHeap{}
	// heap 调用Init初始化堆 之后使用不用主动调整堆
	heap.Init(min)
	heap.Init(max)
	return MedianFinder{max, min}
}

func (this *MedianFinder) AddNum(num int) {
	// num 小于最小堆中的最小值
	if this.maxHeap.Len() == 0 || num > (*this.maxHeap)[0] {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, -num)
	}
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		// pop出最小堆中的最小值
		// 这调的pop不是上面自定义的pop 而是先把第0和n-1个交换 再拿出n-1个
		// 实际是把n-1个元素脱离开这个堆
		// refer to https://studygolang.com/articles/13173#43-heappop
		heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
	} else if this.minHeap.Len() > this.maxHeap.Len()+1 {
		heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() < this.minHeap.Len() {
		return float64(-(*this.minHeap)[0])
	} else if this.maxHeap.Len() > this.minHeap.Len() {
		return float64((*this.maxHeap)[0])
	}
	return float64(-(*this.minHeap)[0]+(*this.maxHeap)[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

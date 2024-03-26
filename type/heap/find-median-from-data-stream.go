/*
 * @Author: hongwei.sun
 * @Date: 2024-03-26 22:16:15
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-26 23:48:21
 * @Description: file content
 */
//  https://leetcode.com/problems/find-median-from-data-stream/
// The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

// For example, for arr = [2,3,4], the median is 3.
// For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
// Implement the MedianFinder class:

// MedianFinder() initializes the MedianFinder object.
// void addNum(int num) adds the integer num from the data stream to the data structure.
// double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.
 

// Example 1:

// Input
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// Output
// [null, null, null, 1.5, null, 2.0]

// Explanation
// MedianFinder medianFinder = new MedianFinder();
// medianFinder.addNum(1);    // arr = [1]
// medianFinder.addNum(2);    // arr = [1, 2]
// medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
// medianFinder.addNum(3);    // arr[1, 2, 3]
// medianFinder.findMedian(); // return 2.0
 

// Constraints:

// -105 <= num <= 105
// There will be at least one element in the data structure before calling findMedian.
// At most 5 * 104 calls will be made to addNum and findMedian.
 

// Follow up:

// If all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
// If 99% of all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
/*
1, maintain two heap, one maxheap, one minheap, and all values in maxheap are not bigger than all values in minheap
and maxheap len approximately minheap len (not bigger than 1) and maxHeap should be bigger than minheap len or equal
in this case, we are able to find median value
2, when one number coming, we need to compare max value in maxheap and min value in minheap
3, then switch them if max vaule from maxheap is bigger than min value in minheap
4, then find median just by max vaule in maxheap and min value in minheap.
*/
type MedianFinder struct {
    minHeap *minHeap
	maxHeap *maxHeap
}


func Constructor() MedianFinder {
	// 注意这里要取地址 因为我们定义的是个指针 
	// 为什么必须定义指针是因为pop和push的方法必须得用指针
    minHeap := &minHeap{}
	maxHeap := &maxHeap{}
	// 这里好像也可以不用init
	heap.Init(minHeap)
	heap.Init(maxHeap)
	// 注意结构体这要有逗号 否则有语法错误
	return MedianFinder {
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
}


func (this *MedianFinder) AddNum(num int)  {
	// 这个方法是关键 要实现我们上面说的1,2,3
	// 注意这里调用的是heap 这个类 不是this.maxHeap, this.maxHeap 是参数
	// 先都往maxheap里push 然后把最大值放到minheap里再pop this maxheap
	// 这样能保证maxheap和minheap长度近似相等
	heap.Push(this.maxHeap, num)
	// 这里要用指针 因为我们定义的maxHeap是个指针，现在要拿这个maxheap的最大值
	heap.Push(this.minHeap, (*this.maxHeap)[0])
	heap.Pop(this.maxHeap)
	// 	maxHeap should be bigger than minheap len or equal
	//  in this case, we are able to find median value
	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(this.maxHeap, (*this.minHeap)[0])
		heap.Pop(this.minHeap)
		
	}
}


func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() == this.minHeap.Len() {
		return float64((*this.maxHeap)[0]+(*this.minHeap)[0])/2
	}
	return float64((*this.maxHeap)[0])
}

type maxHeap []int
func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Pop() interface{} {
	num := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return num
}
func (h *maxHeap) Push(num interface{}) {
	*h = append(*h, num.(int))
}


type minHeap []int
func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Pop() interface{} {
	num := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return num
}
func (h *minHeap) Push(num interface{}) {
	*h = append(*h, num.(int))
}
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
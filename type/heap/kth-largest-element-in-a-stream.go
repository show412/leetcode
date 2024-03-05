/*
 * @Author: hongwei.sun
 * @Date: 2024-03-05 21:40:00
 * @LastEditors: your name
 * @LastEditTime: 2024-03-05 21:40:00
 * @Description: file content
 */
import (
	"container/heap"
)

/*
 * @Author: hongwei.sun
 * @Date: 2024-02-16 15:49:42
 * @LastEditors: your name
 * @LastEditTime: 2024-03-05 20:00:52
 * @Description: file content
 */
// https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
// Design a class to find the kth largest element in a stream.
// Note that it is the kth largest element in the sorted order,
// not the kth distinct element.

// Implement KthLargest class:

// KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
// int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.

// Example 1:

// Input
// ["KthLargest", "add", "add", "add", "add", "add"]
// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
// Output
// [null, 4, 5, 5, 8, 8]

// Explanation
// KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
// kthLargest.add(3);   // return 4
// kthLargest.add(5);   // return 5
// kthLargest.add(10);  // return 5
// kthLargest.add(9);   // return 8
// kthLargest.add(4);   // return 8

// Constraints:

// 1 <= k <= 104
// 0 <= nums.length <= 104
// -104 <= nums[i] <= 104
// -104 <= val <= 104
// At most 104 calls will be made to add.
// It is guaranteed that there will be at least k elements in the array when you search for the kth element.

type KthLargest struct {
	minHeap *minHeap
	K       int
}

func Constructor(k int, nums []int) KthLargest {
	h := &minHeap{}
	heap.Init(h)
	kthObj := KthLargest{minHeap: h, K: k}
	for _, num := range nums {
		kthObj.Add(num)
	}
	return kthObj
}

func (this *KthLargest) Add(val int) int {
	// notice call push use heap to call
	// because it's heap function, we define below is only part of the two function called at last
	heap.Push(this.minHeap, val)
	if this.minHeap.Len() > this.K {
		heap.Pop(this.minHeap)
	}
	// notice there should be one pointer
	return (*this.minHeap)[0]
}

type minHeap []int

// these len less swap function is for Push and Pop function in heap
func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *minHeap) Push(x interface{}) {
	// this is not heap push function, only tell heap function how to add one item in heap
	/*
		we can see in real heap push function, there is one up adjust
		func Push(h Interface, x any) {
			h.Push(x)
			up(h, h.Len()-1)
		}

	*/
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	// here actually in golang container heap, it swap 0 index of heap and last index of heap
	// and down adjust this heap then return this last one
	/*
		this is inner container heap pop
		we can see it call customized pop fuction at last
		why we need to define it, because perhaps it's not an array
		func Pop(h Interface) any {
			n := h.Len() - 1
			h.Swap(0, n)
			down(h, 0, n)
			return h.Pop()
		}
	*/
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
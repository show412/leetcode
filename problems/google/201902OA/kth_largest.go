import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/
// refer to https://www.jianshu.com/p/c9664734033e
// 可以直接排序 O(NlogN) O(1)
// 可以构建最小堆排序 也可以构建最大堆排序 refer to https://blog.csdn.net/mofiu/article/details/83620743
// O(Nlogk) O(k)
// 也可以用 quickSelect 算法 O(N) in the average case, {O}(N^2)in the worst case.
// space complexity is O(1)

// heap by myself, it has some wrong, couldn't have the right result
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// create one minHeap including k numbers in nums
	for i := 0; i <= (k-1)/2; i++ {
		minHeap(i, k-1, nums)
	}
	fmt.Println(nums)
	// from k to len-1, compare the num with top num of heap and adjust heap
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[i], nums[0] = nums[0], nums[i]
			minHeap((k-2)/2, k-1, nums)
		}
	}
	fmt.Println(nums)
	return nums[0]
}

func minHeap(root int, end int, nums []int) {
	child := 2*root + 1
	for {
		if child > end {
			break
		}
		if (child+1) <= end && nums[child] > nums[child+1] {
			child += 1
		}
		if nums[root] > nums[child] {
			nums[child], nums[root] = nums[root], nums[child]
			// recusive to adjust the next tree into minHeap
			root = child
		} else {
			break
		}
	}
}

// golang has the heap struct as sort
// implement the interface as these function (sort.Interface and push, pop of heap)
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func findKthLargest(nums []int, k int) int {
	pq := &IntHeap{}
	for i, v := range nums {
		if i < k {
			heap.Push(pq, v)
		} else {
			if v > (*pq)[0] {
				heap.Pop(pq)
				heap.Push(pq, v)
			}
		}
	}

	return (*pq)[0]
}

// quick select and partition
// refer to https://leetcode.com/problems/kth-largest-element-in-an-array/solution/
func findKthLargest(nums []int, k int) int {

}

// sort O(NlogN) O(1)
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[k-1]
}

/*
 * @Author: hongwei.sun
 * @Date: 2024-03-05 21:40:27
 * @LastEditors: your name
 * @LastEditTime: 2024-03-05 21:57:00
 * @Description: file content
 */
import "container/heap"

//  https://leetcode.com/problems/last-stone-weight/
//  You are given an array of integers stones where stones[i] is the weight of the ith stone.

//  We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:

//  If x == y, both stones are destroyed, and
//  If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
//  At the end of the game, there is at most one stone left.

//  Return the weight of the last remaining stone. If there are no stones left, return 0.

//  Example 1:

//  Input: stones = [2,7,4,1,8,1]
//  Output: 1
//  Explanation:
//  We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
//  we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
//  we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
//  we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
//  Example 2:

//  Input: stones = [1]
//  Output: 1

//  Constraints:

//  1 <= stones.length <= 30
//  1 <= stones[i] <= 1000
// max heap in golang

func lastStoneWeight(stones []int) int {
	maxHeap := &maxHeap{}
	heap.Init(maxHeap)
	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}
	for maxHeap.Len() > 1 {
		s1 := heap.Pop(maxHeap).(int)
		s2 := heap.Pop(maxHeap).(int)
		if s1 != s2 {
			heap.Push(maxHeap, s1-s2)
		}
	}
	if maxHeap.Len() == 1 {
		return (*maxHeap)[0]
	}
	return 0
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	// here define one max heap by changing Less
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
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

func (h *maxHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}
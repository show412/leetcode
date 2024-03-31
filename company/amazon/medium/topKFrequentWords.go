import "container/heap"

// https://leetcode.com/problems/top-k-frequent-words/
/*
Given a non-empty list of words, return the k most frequent elements.

Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

Example 1:
Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
Output: ["i", "love"]
Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.
Example 2:
Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
Output: ["the", "is", "sunny", "day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
    with the number of occurrence being 4, 3, 2 and 1 respectively.
Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Input words contain only lowercase letters.
Follow up:
Try to solve it in O(n log k) time and O(n) extra space.
*/
type item struct {
	str   string
	count int
}
type ItemHeap []item

func (h ItemHeap) Len() int { return len(h) }
func (h ItemHeap) Less(i, j int) bool {
	if h[i].count > h[j].count {
		return true
	}
	if h[i].count == h[j].count {
		return (h[i].str < h[j].str)
	}
	return false
}
func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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

func topKFrequent(words []string, k int) []string {
	heapItem := &ItemHeap{}
	hashItem := make(map[string]int, 0)
	res := make([]string, 0)
	// 不一定非要init
	heap.Init(heapItem)
	for i := 0; i < len(words); i++ {
		word := words[i]
		if _, ok := hashItem[word]; ok {
			hashItem[word] += 1
		} else {
			hashItem[word] = 1
		}
	}
	for key, value := range hashItem {
		// 修改 less 函数可以变成最大堆
		h := item{key, value}
		heap.Push(heapItem, item(h))
	}
	count := 0
	for count < k {
		// 断言 断言只是在接口和类型之间的转换 这里返回的是interface
		res = append(res, heap.Pop(heapItem).(item).str)
		if len(*heapItem) == 0 {
			break
		}
		count++
	}
	return res

}

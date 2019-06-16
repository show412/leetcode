package main

import (
	"fmt"
)

func main() {
	res := []int{3, 5, 6, 1, 2, 8}
	HeapSort(res)
	fmt.Println(res)
}

func minHeap(root int, end int, c []int) {
	for {
		var child = 2*root + 1
		//判断是否存在child节点
		if child > end {
			break
		}
		//判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}

//降序排序
func HeapSort(c []int) {
	var n = len(c) - 1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, c)
	}
	fmt.Println("堆构建完成")
	for end := n; end >= 0; end-- {
		if c[0] < c[end] {
			c[0], c[end] = c[end], c[0]
			minHeap(0, end-1, c)
		}
	}
}

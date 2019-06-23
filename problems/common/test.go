package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	// [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(res)
}

// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
func dailyTemperatures(T []int) []int {
	if len(T) == 1 {
		return []int{0}
	}
	var res []int
	type node struct {
		value  int
		index  int
		bigger int
	}
	var stack []*node
	var nodeList []*node
	for i := 0; i < len(T); i++ {
		item := node{value: T[i], index: i, bigger: 0}
		nodeList = append(nodeList, &item)
	}
	stack = append(stack, nodeList[0])
	for i := 1; i < len(nodeList); i++ {
		for len(stack) != 0 && nodeList[i].value > stack[len(stack)-1].value {
			last := stack[len(stack)-1]
			last.bigger = i - last.index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nodeList[i])
	}
	for i := 0; i < len(nodeList); i++ {
		res = append(res, nodeList[i].bigger)
	}
	return res
}

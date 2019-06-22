package main

import (
	"fmt"
	"strconv"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := summaryRanges([]int{0})
	fmt.Println(res)
}

// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
func summaryRanges(nums []int) []string {
	// m := make(map[int]bool, 0)
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	var res []string
	start := 1
	end := 0
	last := 0
	for start < len(nums) {
		if nums[start] != nums[last]+1 {
			if last-end >= 1 {
				ranges := strconv.Itoa(nums[end]) + "->" + strconv.Itoa(nums[last])
				res = append(res, ranges)
			} else {
				res = append(res, strconv.Itoa(nums[end]))
			}
			end = start
		}
		last = start
		start++
	}
	if end == len(nums)-1 {
		res = append(res, strconv.Itoa(nums[end]))
	} else {
		ranges := strconv.Itoa(nums[end]) + "->" + strconv.Itoa(nums[last])
		res = append(res, ranges)
	}
	return res
}

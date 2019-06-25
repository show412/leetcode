package main

import (
	"fmt"
	"strings"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := lengthLongestPath("a\n\tb.txt\na2\n\tb2.txt")
	// [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(res)
}

func lengthLongestPath(input string) int {
	if input == "" {
		return 0
	}
	inputArray := strings.Split(input, "\n")
	res := 0
	// 存的是栈里从当前到栈底的字符串长度总数
	var stack []int
	// insert dummy length
	// stack = append(stack, 0)
	for i := 0; i < len(inputArray); i++ {
		// the level remove the '\t' so +1
		level := strings.LastIndexByte(inputArray[i], '\t')
		// because there is a 0 in the bottom of stack so +1
		for level+1 < len(stack) {
			stack = stack[:len(stack)-1]
		}
		// because there is a '/' at the end so +1
		if len(stack) == 0 {
			stack = append(stack, len(inputArray[i]))
		} else {
			stack = append(stack, stack[len(stack)-1]+len(inputArray[i][level:]))
		}

		// for the last file, remove the '/' count, so  -1
		if strings.LastIndexByte(inputArray[i], '.') >= 0 {
			res = max(res, stack[len(stack)-1])
		}
	}
	return res
}

// func lengthLongestPath(input string) int {
// 	if input == "" {
// 		return 0
// 	}
// 	inputArray := strings.Split(input, "\n")
// 	res := 0
// 	// 存的是栈里从当前到栈底的字符串长度总数
// 	var stack []int
// 	// insert dummy length
// 	stack = append(stack, 0)
// 	for i := 0; i < len(inputArray); i++ {
// 		// the level remove the '\t' so +1
// 		level := strings.LastIndexByte(inputArray[i], '\t') + 1
// 		fmt.Println(level)
// 		// because there is a 0 in the bottom of stack so +1
// 		for level+1 < len(stack) {
// 			stack = stack[:len(stack)-1]
// 		}
// 		// because there is a '/' at the end so +1
// 		stack = append(stack, stack[len(stack)-1]+len(inputArray[i][level:])+1)
// 		// for the last file, remove the '/' count, so  -1
// 		if strings.LastIndexByte(inputArray[i], '.') >= 0 {
// 			res = max(res, stack[len(stack)-1]-1)
// 		}
// 	}
// 	return res
// }

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

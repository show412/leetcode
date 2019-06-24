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
	if strings.IndexByte(input, '.') >= 0 && strings.LastIndexByte(input, '\t') <= -1 {
		return len(input)
	}
	inputArray := strings.Split(input, "\n")
	res := len(inputArray[0])

	for start := 1; start < len(inputArray); start++ {
		cur := inputArray[start]

		end := start + 1
		curIndex := strings.LastIndexByte(cur, '\t')
		length := len(cur)
		if curIndex >= 0 {
			length = len(cur[curIndex:])
		}

		if start == 1 && strings.LastIndexByte(cur, '.') >= 0 {
			res = max(res, res+length)
			// fmt.Println("cur", res)
			continue
		}
		if end < len(inputArray) && cur[0] == '\t' && cur[1] != '\t' {
			// fmt.Println("********")
			// // fmt.Println(start)
			// fmt.Println(cur)
			// fmt.Println(len(cur[curIndex:]))

			for end < len(inputArray) && (end-start) < len(inputArray[end]) && inputArray[end][end-start] == '\t' {

				next := inputArray[end]
				// fmt.Println(next)

				nextIndex := strings.LastIndexByte(next, '\t')
				if nextIndex >= 0 {
					length += len(next[nextIndex:])
				} else {
					length += len(next)
				}

				// fmt.Println(len(next[nextIndex:]))
				if strings.IndexByte(next, '.') >= 0 {
					res = max(res, len(inputArray[0])+length)
					// fmt.Println("next", res)
					break
				}
				end++
			}
		} else {
			start = end
			continue
		}
	}
	if res == len(inputArray[0]) {
		return 0
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

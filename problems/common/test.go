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
	res := lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext")
	// [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(res)
}

func lengthLongestPath(input string) int {
	if input == "" {
		return 0
	}
	inputArray := strings.Split(input, "\n")
	res := len(inputArray[0])
	// ["dir" "\tsubdir1" "\tsubdir2" "\t\tfile.ext"]

	for start := 1; start < len(inputArray); start++ {
		cur := inputArray[start]
		length := len(cur)
		end := start + 1
		if end < len(inputArray) && cur[0] == '\t' && cur[1] != '\t' {
			fmt.Println(start)
			if strings.IndexByte(cur, '.') >= 0 {
				res = max(res, res+len(cur))
				continue
			}
			next := inputArray[end]

			for end < len(inputArray) && (end-start) < len(next) && next[end-start] == '\t' {
				length += len(next)
				if strings.IndexByte(next, '.') >= 0 {
					res = max(res, res+length-end+start)
					break
				}
				end++
			}
		} else {
			start = end + 1
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

package main

import (
	"fmt"
	// "math"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	// 1, 2, 3, 0, 2
	// output [[1,1,0],[1,0,1],[0,0,0]]
	// expected [[1,0,0],[0,1,0],[1,1,1]]
	res := beautifulArray(4)
	fmt.Println(res)
}

func beautifulArray(N int) []int {
	res := make([]int, 1)
	res[0] = 1
	for len(res) < N {
		var tmp []int
		for i := 0; i < len(res); i++ {
			if (2*res[i] - 1) <= N {
				tmp = append(tmp, 2*res[i]-1)
			}
		}
		for i := 0; i < len(res); i++ {
			if 2*res[i] <= N {
				tmp = append(tmp, 2*res[i])
			}
		}
		fmt.Println(tmp)
		cpy := make([]int, len(tmp))
		copy(cpy, tmp)
		res = cpy
	}
	return res
}

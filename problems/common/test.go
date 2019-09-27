package main

import (
	"fmt"
	"sort"
)

/*
test case:
101, 56, 69, 48, 30
4

[8,85,24,85,69]
4

[73,106,39,6,26,15,30,100,71,35,46,112,6,60,110]
29

[98,60,24,89,84,51,61,96,108,87,68,29,14,11,13,50,13,104,57,8,57,111,92,87,9,59,65,116,56,39,55,11,21,105,57,36,48,93,20,94,35,68,64,41,37,11,50,47,8,9]
439
*/
func main() {
	res := numFriendRequests([]int{98, 60, 24, 89, 84, 51, 61, 96, 108, 87, 68, 29, 14, 11, 13, 50, 13, 104, 57, 8, 57, 111, 92, 87, 9, 59, 65, 116, 56, 39, 55, 11, 21, 105, 57, 36, 48, 93, 20, 94, 35, 68, 64, 41, 37, 11, 50, 47, 8, 9})
	// res := canFinish(4, [][]int{[]int{0, 1}})
	fmt.Println(res)
}

// type m map[int]int

func numFriendRequests(ages []int) int {
	if len(ages) == 1 {
		return 0
	}
	res := 0
	sort.Ints(ages)
	fmt.Println(ages)
	for i := len(ages) - 1; i > 0; i-- {
		age := ages[i]
		if age >= 100 {
			// age[B] <= 0.5 * age[A] + 7
			j := sort.Search(len(ages[0:i]), func(j int) bool { return ages[j] > (age/2 + 7) })
			fmt.Println("**>100**")
			fmt.Println(i)
			fmt.Println(j)
			// if j < i {
			a := i - j
			res += a
			// }

		} else {
			j := sort.Search(len(ages[0:i]), func(j int) bool { return ages[j] > (age/2+7) && ages[j] <= 100 })
			fmt.Println("**<100**")
			fmt.Println(i)
			fmt.Println(j)
			// if j < i {
			a := i - j
			res += a
			// }
		}

	}

	m := make(map[int]int, 0)
	for i := 0; i < len(ages); i++ {
		age := ages[i]
		if _, ok := m[age]; !ok {
			m[age] = 1
		} else {
			fmt.Println("add")
			if age > 14 {
				res += m[age]
			}
			m[age] += 1
		}
	}
	return res
}

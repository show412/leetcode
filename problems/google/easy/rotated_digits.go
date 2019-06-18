// https://leetcode.com/problems/rotated-digits/solution/
/*
X is a good number if after rotating each digit individually by 180 degrees,
we get a valid number that is different from X.
Each digit must be rotated - we cannot choose to leave it alone.

A number is valid if each digit remains a digit after rotation.
0, 1, and 8 rotate to themselves;
2 and 5 rotate to each other;
6 and 9 rotate to each other,
and the rest of the numbers do not rotate to any other number and become invalid.

Now given a positive number N, how many numbers X from 1 to N are good?

Example:
Input: 10
Output: 4
Explanation:
There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.
*/
// it's a wrong solution
/* func rotatedDigits(N int) int {
	if N < 0 {
		return 0
	}
	res := make([]int, 0, N)
	sub := make([]string, 0, N)
	// visit := make(map[int]bool, 0)
	nums := []int{0, 1, 2, 5, 6, 8, 9}
	dfs(nums, N, sub, &res)
	return len(res)
}

func dfs(nums []int, max int, sub []string, res *[]int) {
	subString := strings.Join(sub, "")
	if strings.IndexAny(subString, "2569") >= 0 {
		v, _ := strconv.Atoi(subString)
		if v != 0 && v <= max {
			// fmt.Println(v)
			*res = append(*res, v)
		} else {
			return
		}
	}

	// if v > max {
	// 	return
	// }

	for i := 0; i < len(nums); i++ {
		// if len(*res) == 0 && len(visit) == 0 && i == 0 {
		// 	continue
		// }
		// if visit[nums[i]] == true {
		// 	continue
		// } else {
		sub = append(sub, strconv.Itoa(nums[i]))
		// visit[nums[i]] = true
		v, _ := strconv.Atoi(strings.Join(sub, ""))
		if v > max {
			// sub = sub[:len(sub)-1]
			// visit[nums[i]] = false
			break
		}
		dfs(nums, max, sub, res)
		sub = sub[:len(sub)-1]
		// visit[nums[i]] = false
		// }
	}
	return
}
*/
// TC is O(NlogN) The SC is O(logN) if we consider the good function store
func rotatedDigits(N int) int {
	if N < 0 {
		return 0
	}
	cnt := 0
	for i := 1; i <= N; i++ {
		if good(i) == true {
			// fmt.Println(i)
			cnt++
		}
	}
	return cnt
}
func good(num int) bool {
	flag := 0
	for num != 0 {
		if num%10 == 3 || num%10 == 4 || num%10 == 7 {
			return false
		}
		if num%10 == 0 || num%10 == 1 || num%10 == 8 {
			num = num / 10
			continue
		}
		if num%10 == 2 || num%10 == 5 || num%10 == 6 || num%10 == 9 {
			num = num / 10
			flag = 1
			continue
		}
	}
	if flag == 1 {
		return true
	}
	return false
}

// It has a DP solution wiht TC O(logN), we need to investigate

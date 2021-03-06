import "sort"

// https://leetcode.com/problems/friends-of-appropriate-ages/
/*
Some people will make friend requests.
The list of their ages is given and ages[i] is the age of the ith person.

Person A will NOT friend request person B (B != A)
if any of the following conditions are true:

age[B] <= 0.5 * age[A] + 7
age[B] > age[A]
age[B] > 100 && age[A] < 100
Otherwise, A will friend request B.

Note that if A requests B, B does not necessarily request A.
Also, people will not friend request themselves.

How many total friend requests are made?

Example 1:

Input: [16,16]
Output: 2
Explanation: 2 people friend request each other.
Example 2:

Input: [16,17,18]
Output: 2
Explanation: Friend requests are made 17 -> 16, 18 -> 17.
Example 3:

Input: [20,30,100,110,120]
Output:
Explanation: Friend requests are made 110 -> 100, 120 -> 110, 120 -> 100.


Notes:

1 <= ages.length <= 20000.
1 <= ages[i] <= 120.
*/
/*
三个条件:
1, 排序 大的只能向小的发 request
2, 二分搜索 分成 >=100 <100 两组
3, 还要向上搜索相同的值并且满足条件 用hash记出现的次数
TC: O(nlogn) + O(nlogn), SC: O(A)  A is the number of ages
*/
func numFriendRequests(ages []int) int {
	if len(ages) == 1 {
		return 0
	}
	res := 0
	sort.Ints(ages)
	for i := len(ages) - 1; i > 0; i-- {
		age := ages[i]
		if age >= 100 {
			j := sort.Search(len(ages[0:i]), func(j int) bool { return ages[j] > (age/2 + 7) })
			if j < i {
				res += i - j
			}

		} else {
			j := sort.Search(len(ages[0:i]), func(j int) bool { return ages[j] > (age/2+7) && ages[j] <= 100 })
			if j < i {
				res += i - j
			}
		}

	}
	m := make(map[int]int, 0)
	for i := 0; i < len(ages); i++ {
		age := ages[i]
		if _, ok := m[age]; !ok {
			m[age] = 1
		} else {
			if age > 14 {
				res += m[age]
			}
			m[age] += 1
		}
	}
	return res
}

// TC is O(A^2 + N) A is 121; SC is O(A)
func numFriendRequests(ages []int) int {
	if len(ages) == 1 {
		return 0
	}
	res := 0
	count := make([]int, 121)
	for _, age := range ages {
		count[age]++
	}
	for ageA := 0; ageA < 121; ageA++ {
		countA := count[ageA]
		for ageB := 0; ageB < 121; ageB++ {
			countB := count[ageB]
			if ageA/2+7 >= ageB {
				continue
			}
			if ageA < ageB {
				continue
			}
			if ageA < 100 && 100 < ageB {
				continue
			}
			/*
				f ageA == ageB, then we overcounted: as you cannot friend request yourself.
			*/
			res += countA * countB
			if ageA == ageB {
				res -= countA
			}
		}
	}
	return res
}

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
1, 通过 hash key 是 age value 是公式值
2, 排序 大的只能向小的发 request
3, 二分 分成 >100 <=100 两组
*/
func numFriendRequests(ages []int) int {
	if len(ages) == 1 {
		return 0
	}

}

import "math"

// https://leetcode.com/problems/maximize-distance-to-closest-person/
/*
In a row of seats, 1 represents a person sitting in that seat, and 0 represents that the seat is empty.

There is at least one empty seat, and at least one person sitting.

Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return that maximum distance to closest person.

Example 1:

Input: [1,0,0,0,1,0,1]
Output: 2
Explanation:
If Alex sits in the second open seat (seats[2]), then the closest person has distance 2.
If Alex sits in any other open seat, the closest person has distance 1.
Thus, the maximum distance to the closest person is 2.
Example 2:

Input: [1,0,0,0]
Output: 3
Explanation:
If Alex sits in the last seat, the closest person is 3 seats away.
This is the maximum distance possible, so the answer is 3.
Note:

1 <= seats.length <= 20000
seats contains only 0s or 1s, at least one 0, and at least one 1.
*/
/*
Next Array solution
Time Complexity: O(N), where NN is the length of seats.
Space Complexity: O(N), the space used by left and right.
there are two other solution with SC O(1) solution
https://leetcode.com/problems/maximize-distance-to-closest-person/solution/
*/
func maxDistToClosest(seats []int) int {
	// the varaible could be difined as left and right
	Forward := make([]int, len(seats))
	Afterward := make([]int, len(seats))
	curF := -1
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			curF = i
			Forward[i] = 0
			continue
		}
		if curF == -1 {
			Forward[i] = math.MaxInt32
		} else {
			Forward[i] = i - curF
		}
	}
	curA := -1
	for i := len(seats) - 1; i >= 0; i-- {
		if seats[i] == 1 {
			curA = i
			Afterward[i] = 0
			continue
		}
		if curA == -1 {
			Afterward[i] = math.MaxInt32
		} else {
			Afterward[i] = curA - i
		}
	}
	res := 0
	for i := 0; i < len(seats); i++ {
		res = max(res, min(Forward[i], Afterward[i]))
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/*
class Solution {
    public int maxDistToClosest(int[] seats) {
        int N = seats.length;
        int K = 0; //current longest group of empty seats
        int ans = 0;

        for (int i = 0; i < N; ++i) {
            if (seats[i] == 1) {
                K = 0;
            } else {
                K++;
                ans = Math.max(ans, (K + 1) / 2);
            }
        }

        for (int i = 0; i < N; ++i)  if (seats[i] == 1) {
            ans = Math.max(ans, i);
            break;
        }

        for (int i = N-1; i >= 0; --i)  if (seats[i] == 1) {
            ans = Math.max(ans, N - 1 - i);
            break;
        }

        return ans;
    }
}
*/

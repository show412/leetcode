/**
 * @param arrayIn: The original array.
 * @return: Count the minimum number of subarrays.
 */
//TC is O(n^2) with DP. it's the right solution
// for TC is O(nlogn), refer to LIS.go
func LeastSubsequences(arrayIn []int) int {
	if len(arrayIn) == 0{
		return 1
	}
	f := make([]int, len(arrayIn))
	f[0] = 1
	res := 1
	for i := 1; i < len(arrayIn); i++ {
		f[i] = 1
		for j :=0; j<i; j++ {
			if arrayIn[i] < arrayIn[j] {
				f[i] = int(math.Max(float64(f[i]),float64(f[j]+1)))
			}
		}
		res = int(math.Max(float64(f[i]),float64(res)))
	}
	return res
}

func LeastSubsequences(arrayIn []int) int {
	// Write your code here.
	if len(arrayIn) == 0 {
		return 1
	}
	m := make(map[int]int)
	for i := 0; i < len(arrayIn); i++ {
		m[arrayIn[i]] = i
	}
	arrayIn = quick_sort(arrayIn, 0, len(arrayIn)-1)
	// sort.Ints(arrayIn)
	res := 1
	for i := len(arrayIn) - 1; i > 0; i-- {
		if m[arrayIn[i]] < m[arrayIn[i-1]] {
			continue
		} else {
			res++
		}
	}
	return res
}

func quick_sort(nums []int, l int, h int) []int {
	left := l
	right := h
	if left >= right {
		return nil
	}
	pivot := nums[l]
	for l != h {
		// notice !! it must that the right go firstly
		for l < h && nums[h] >= pivot {
			h--
		}
		for l < h && nums[l] <= pivot {
			l++
		}
		if l < h {
			nums[l], nums[h] = nums[h], nums[l]
		}
	}

	nums[left] = nums[l]
	nums[l] = pivot
	quick_sort(nums, left, l-1)
	quick_sort(nums, l+1, right)
	return nums
}

public class Solution {
    /**
     * @param arrayIn: The original array.
     * @return: Count the minimum number of subarrays.
     */
    public int LeastSubsequences(List<Integer> arrayIn) {
        // Write your code here.
        int n = arrayIn.size();
        int[] dp = new int[n];
        dp[0] = 1;
        int ans = 1;
        for(int i = 1; i < n ;i++) {
            dp[i] = 1;
            for(int j = 0; j < i; j++) {
                if(arrayIn.get(i) >= arrayIn.get(j)) {
                    dp[i] = Math.max(dp[i], dp[j]+1);
                }
            }
            ans = Math.max(ans, dp[i]);
        }
        return ans;
    }
}
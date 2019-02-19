// https://leetcode.com/problems/next-permutation/submissions/
// 按字典序升序排列的函数 输入的是一个数组  114532  输出应该是115234

class Solution {
  public void swapItem(int[] nums, int i, int j) {
    	int temp = nums[i];
		nums[i] = nums[j];
		nums[j] = temp;
	}
	public void swapList(int[] nums, int i, int j) {
		while (i < j) {
			swapItem(nums, i, j);
			i ++; j --;
		}
  }
  
    public int[] nextPermutation(int[] nums) {
      int len = nums.length;
      if ( len <= 1)
        return nums;
      int i = len - 1;
      // 114532  532是一个降序排列 不论如何调整都不会有满足条件的下一个字典序结果
      // 所以先找到一个降序子序列
      while (i > 0 && nums[i] <= nums[i - 1])
        i --;
        // 使这个降序子序列变为升序子序列
      swapList(nums, i, len - 1);
      if (i != 0) {
        int j = i;
        // 把4用这个升序子序列里大于它的最小值交换 就是最后的结果
        while (nums[j] <= nums[i - 1]) j++;
        swapItem(nums, j, i-1);
      }
      return nums;
    }
}
# @param {Integer[]} nums1
# @param {Integer[]} nums2
# @return {Float}
#     对于一个长度为n的已排序数列a，若n为奇数，中位数为第(n/2+1)个数 ,
#     若n为偶数，则中位数=[第(n/2)个数 + 第(n/2+1)个数] / 2
#     如果我们可以在两个数列中求出第K小的元素，便可以解决该问题
#     不妨设数列A元素个数为n，数列B元素个数为m，各自升序排序，求第k小元素
#     取A[k / 2] B[k / 2] 比较，
#     如果 A[k / 2] > B[k / 2] 那么，所求的元素必然不在B的前k / 2个元素中(证明反证法)
#     反之，必然不在A的前k / 2个元素中，于是我们可以将A或B数列的前k / 2元素删去，求剩下两个数列的
#     k - k / 2小元素，于是得到了数据规模变小的同类问题，递归解决
#     如果 k / 2 大于某数列个数，所求元素必然不在另一数列的前k / 2个元素中，同上操作就好。

def find_median_sorted_arrays(nums1, nums2)
  m = nums1.length
  n = nums2.length
  sum = m + n
  if (sum%2 == 1)
    result = find_kth(nums1,nums2,sum/2 + 1)
  else
    result = (find_kth(nums1,nums2,sum/2) + find_kth(nums1,nums2,sum/2 + 1))/2.to_f
  end
  return result
end

def find_kth(array1, array2,k)
  return array2[k-1] if array1.length == 0 || array1 == nil || array1[0] == nil
  return array1[k-1] if array2.length == 0 || array2 == nil || array2[0] == nil
  return [array1[0], array2[0]].min if k == 1
  array1_key = k/2 - 1 > array1.length-1  ? 4294967296 : array1[k/2 - 1]
  array2_key = k/2 - 1 > array2.length-1 ? 4294967296 : array2[k/2 - 1]
  if(array1_key > array2_key)
    array2_slice = array2.slice(k/2, array2.length-1)
    return find_kth(array1, array2_slice, k-k/2)
  else
    array1_slice = array1.slice(k/2, array1.length-1)
    return find_kth(array1_slice, array2, k-k/2)
  end
end

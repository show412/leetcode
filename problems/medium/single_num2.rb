# https://leetcode.com/problems/single-number-ii/
# 因为其他数是出现三次的，也就是说，对于每一个二进制位，如果只出现一次的数在该二进制位为1，那么这个二进制位在全部数字中出现次数无法被3整除。
# 对于每一位，我们让Two，One表示当前位的状态。

# 我们看Two和One里面的每一位的定义，对于ith(表示第i位)：

# 如果Two里面ith是1，则ith当前为止出现1的次数模3的结果是2
# 如果One里面ith是1，则ith目前为止出现1的次数模3的结果是1
# 注意Two和One里面不可能ith同时为1，因为这样就是3次，每出现3次我们就可以抹去（消去）。
# 那么最后One里面存储的就是每一位模3是1的那些位，综合起来One也就是最后我们要的结果。

# 如果B表示输入数字的对应位，Two+和One+表示更新后的状态
# 那么新来的一个数B，此时跟原来出现1次的位做一个异或运算，&上~Two的结果(也就是不是出现2次的)，那么剩余的就是当前状态是1的结果。
# 同理Two ^ B （2次加1次是3次，也就是Two里面ith是1，B里面ith也是1，那么ith应该是出现了3次，
# 此时就可以消去，设置为0），我们相当于会消去出现3次的位。

# 但是Two ^ B也可能是ith上Two是0，B的ith是1，这样Two里面就混入了模3是1的那些位！！！
# 怎么办？我们得消去这些！我们只需要保留不是出现模3余1的那些位ith，
# 而One是恰好保留了那些模3余1次数的位，`取反不就是不是模3余1的那些位ith么？最终对(~One+)取一个&即可。

# 也可以这么理解
# 对 ones操作， ones = ones ^ A[i] & (~tows)意思是把A[i] 计入到出现一次的情况里，但是有可能这个数已经出现过两次，所以要排除两次情况
# 对 twos操作， twos = twos ^ A[i] & (~ones) 把A[i] 计入到两次情况，如果第三次出现了，自然消掉了，如果第二次出现，那么就计入，如果是第一次出现，不计入，这个过程通过和ones的非相交来控制



# b = (b^i)&~a;
# a = (a^i)&~b;
# 第一次：a = 0,b= A[i];
# 第二次：a = A[i], b = 0;
# 第三次：a= 0, b = 0;
# 返回b就是只出现一次的数。

# @param {Integer[]} nums
# @return {Integer}
def single_number(nums)
  return 0 if nums.length == 0
  one = 0
  two = -1
  i = 0
  while(i<nums.length)
    # why two should be first, and one is last? to be though more
    two = (two^nums[i])&~one
    one = (one^nums[i])&~two
    i+=1
  end
  return one
end

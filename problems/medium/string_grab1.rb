# def solution(a, b)
#   # write your code in Ruby 2.2
#   return nil if (a == 0 && b > 3) || ( a > 3 && b == 0)
#   target = []
#   (0..a).each do |i|
#     (0..b).each do |j|

#     end
#   end 
#   # a.times { target << "a" }
#   # b.times { target << "b" }
#   # path = []
#   # @result = []
#   # target.sort!
#   # subsetsHelper(path, target, 0, @result)
#   # @result  
# end

# def subsets(nums)
#   path = []
#   @result = []
#   nums.sort!
#   subsetsHelper(path, nums, 0, @result)
#   @result
# end

# def subsetsHelper(path, nums, pos, result)
#   result << path
#   i = pos
#   while(i < nums.length)
#     path << nums[i]
#     subsetsHelper(path, nums, i + 1, result)
#     path.pop
#     i += 1
#   end
# end


# public class Solution {
#     /*
#      * @param str: A string
#      * @return: all permutations
#      */
#     public List<String> stringPermutation2(String str) {
#         // write your code here
#         List<String> result = new ArrayList<>();
        
#         if (str == null || str.length() == 0) {
#             result.add("");
#             return result;
#         }
        
#         //Sort Array beforehand for ordering purpose
#         str = sortString(str);
        
#         Set<String> set = new HashSet<>(); 
#         boolean[] isUsed = new boolean[str.length()]; 
#         String testStr = "";
        
#         // DFS
#         DFS(str, testStr, isUsed,  set, result);
#         return result;
#     }
    
#     private void DFS(String str, 
#                      String testStr,
#                      boolean[] isUsed,
#                      Set<String> set,
#                      List<String> result) {
#         if (testStr.length() == str.length()) {
#             result.add(new String(testStr)); 
#             return;
#         }
        
#         for (int i = 0; i < str.length(); i++) {
#             if (isUsed[i] == true || i != 0 && isUsed[i - 1] == false 
#                 && str.charAt(i) == str.charAt(i - 1)) {
#                 continue;
#             }
            
#             if (set.contains(testStr + str.charAt(i))) {
#                 continue;
#             }
            
#             // 字符串拼接的方法
#             testStr = testStr + str.charAt(i); 
#             set.add(testStr);
#             isUsed[i] = true;
           
#             DFS(str, testStr, isUsed,set, result);
            
#             testStr = testStr.substring(0, testStr.length() - 1); 
#             set.remove(testStr);
#             isUsed[i] = false;
#         }
#     }
    
#     private String sortString(String str) {
#         char[] tempArray = str.toCharArray(); 
#         Arrays.sort(tempArray);
#         return new String(tempArray);
#     }
# }

# 作者：6默默Welsh
# 链接：https://www.jianshu.com/p/ccaabff3033a
# 來源：简书
# 简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。

class Array
  #code from http://www.sharejs.com/codes
  def permutations
    return [self] if size < 2
    perm = []
    each { |e| (self - [e]).permutations.each { |p| perm << ([e] + p) } }
    perm
  end
end

p ["a","a","a","b","b","b"].permutations
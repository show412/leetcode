/*
 * @Author: hongwei.sun
 * @Date: 2024-03-31 11:44:33
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-31 11:45:20
 * @Description: file content
 */

//  描述
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param number int整型 
 * @return int整型
*/
func jumpFloor( number int ) int {
    // write code here
    f := make([]int, number+1)
    f[0] = 1
    f[1] = 1
    for i := 2; i <= number; i++ {
        f[i] = f[i-1]+f[i-2]
    }
    return f[number]
}
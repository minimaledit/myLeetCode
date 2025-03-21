/*
https://leetcode.com/problems/climbing-stairs/description/

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example:
Input: n = 3  Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

func climbStairs(n int) int {
    if n==1{
        return 1
    }
    if n==2{
        return 2
    }
    prev1, prev2 := 2, 1
    for i:=2;i<n;i++{
        prev2, prev1 = prev1, prev1+prev2
    }
    return prev1
}

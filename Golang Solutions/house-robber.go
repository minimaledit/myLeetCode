/*
https://leetcode.com/problems/house-robber/description/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police 
if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example:
Input: nums = [1,2,3,1]  Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	prevprev, prev := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		// curr:=max(prev, prevprev+nums[i])
		// prevprev = prev
		// prev = curr
		prev, prevprev = max(prev, prevprev+nums[i]), prev
	}
	return prev
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}


//a more visual example
func rob(nums []int) int {
	houseCount := len(nums)
	if houseCount == 1 {
		return nums[0]
	}
	if houseCount == 2 {
		return max(nums[0], nums[1])
	}
	if houseCount == 3 {
		return max(nums[0]+nums[2], nums[1])
	}
	dp := make([]int, houseCount)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	dp[2] = max(nums[0]+nums[2], nums[1])
	for i := 3; i < houseCount; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[houseCount-1]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

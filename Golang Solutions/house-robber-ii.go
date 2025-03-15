/*
https://leetcode.com/problems/house-robber-ii/description/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle.
That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected,
and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: nums = [2,3,2]  Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

Example 2:
Input: nums = [1,2,3,1]  Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
*/

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(rob2(nums, 0, n-1), rob2(nums, 1, n))
}
func rob2(nums []int, start, houseCount int) int {
	prevprev, prev := nums[0+start], max(nums[0+start], nums[1+start])
	for i := 2 + start; i < houseCount; i++ {
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

/*
https://leetcode.com/problems/majority-element-ii/description/

Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Input: nums = [3,2,3] Output: [3]
*/
func majorityElement(nums []int) []int {
	n := len(nums)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[nums[i]]++
	}
	border := n / 3
	ans := make([]int, 0)
	for key, value := range m {
		if value > border {
			ans = append(ans, key)
		}
	}
	return ans
}

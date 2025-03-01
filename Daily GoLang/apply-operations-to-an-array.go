/*
https://leetcode.com/problems/apply-operations-to-an-array/description/

You are given a 0-indexed array nums of size n consisting of non-negative integers.

You need to apply n - 1 operations to this array where, in the ith operation (0-indexed), you will apply the following on the ith element of nums:

If nums[i] == nums[i + 1], then multiply nums[i] by 2 and set nums[i + 1] to 0. Otherwise, you skip this operation.
After performing all the operations, shift all the 0's to the end of the array.

For example, the array [1,0,2,0,0,1] after shifting all its 0's to the end, is [1,2,1,0,0,0].
Return the resulting array.

Note that the operations are applied sequentially, not all at once.
*/
func applyOperations(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	var j int
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
			i++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			ans[j] = nums[i]
			j++
		}
	}
	return ans
}

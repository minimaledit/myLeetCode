/*
https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/description/

Given an array nums sorted in non-decreasing order, return the maximum between the number of positive integers and the number of negative integers.
In other words, if the number of positive integers in nums is pos and the number of negative integers is neg, then return the maximum of pos and neg.
Note that 0 is neither positive nor negative.

Example:
Input: nums = [-3,-2,-1,0,0,1,2] Output: 3
Explanation: There are 2 positive integers and 3 negative integers. The maximum count among them is 3.
*/

// O(logN)
func maximumCount(nums []int) int {
    return max(len(nums) - upperBound(nums, 0), lowerBound(nums,0))
}
func LowerBound(arr []int, target int) int{
    left, right, mid := 0, len(arr), 0 
    for left<right{
        mid = (left + right) / 2
        if arr[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
    }
    return left
}
func UpperBound(array []int, target int) int {
    low, high := 0, len(array) 
    for low < high {
        mid := (low + high) / 2
        if array[mid] > target {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return low
}

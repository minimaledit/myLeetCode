/*
https://leetcode.com/problems/longest-nice-subarray/description/

You are given an array nums consisting of positive integers.
We call a subarray of nums nice if the bitwise AND of every pair of elements that are in different positions in the subarray is equal to 0.
Return the length of the longest nice subarray.
A subarray is a contiguous part of an array.
Note that subarrays of length 1 are always considered nice.

Example 1:
Input: nums = [1,3,8,48,10]  
Output: 3
Explanation: The longest nice subarray is [3,8,48]. This subarray satisfies the conditions:
- 3 AND 8 = 0.
- 3 AND 48 = 0.
- 8 AND 48 = 0.
It can be proven that no longer nice subarray can be obtained, so we return 3.

Hint 1.
Use a Sliding Window Approach to go over the subarrays. for a valid subarray.
Hint 2.
Maintain a bitmask representing the bitwise OR of all elements. This mask must satisfy (mask & nextElement) == 0 before adding the next element.
Hint 3.
To remove elements from the left of the sliding window, update the mask by unsetting their bits (using XOR (or subtraction) works here since the subarray has no overlapping bits).
Hint 4.
If the next element causes a conflict, move the left pointer until the conflict is resolved ((mask & nextElement) == 0), then continue adding new elements.
*/

func longestNiceSubarray(nums []int) int {
    left, right, mask, maxLen := 0, 0, 0, 0
    for right < len(nums){
        for (mask & nums[right]) != 0 {
            mask -= nums[left]
            left++
        }
        mask+=nums[right]
        if right - left + 1 > maxLen{
            maxLen = right - left + 1
        }
        right++
    }
    return maxLen
}

/*
https://leetcode.com/problems/zero-array-transformation-ii/description/

You are given an integer array nums of length n and a 2D array queries where queries[i] = [l[i], r[i], val[i]].

Each queries[i] represents the following action on nums:
- Decrement the value at each index in the range [l[i], r[i]] in nums by at most vali.
- The amount by which each value is decremented can be chosen independently for each index.
A Zero Array is an array with all its elements equal to 0.

Return the minimum possible non-negative value of k, such that after processing the first k queries in sequence, nums becomes a Zero Array. If no such k exists, return -1.

Input: nums = [2,0,2], queries = [[0,2,1],[0,2,1],[1,1,3]]  Output: 2

Explanation:
For i = 0 (l = 0, r = 2, val = 1):
Decrement values at indices [0, 1, 2] by [1, 0, 1] respectively.
The array will become [1, 0, 1].
For i = 1 (l = 0, r = 2, val = 1):
Decrement values at indices [0, 1, 2] by [1, 0, 1] respectively.
The array will become [0, 0, 0], which is a Zero Array. Therefore, the minimum value of k is 2.
*/
func minZeroArray(nums []int, queries [][]int) int {
    n := len(nums)
    counts := make([]int, n+1)
    k, sum := 0, 0

    for i := 0; i < n; i++ {
        num := nums[i]
        for sum+counts[i] < num {
            if k == len(queries) {
                return -1
            }
            left, right, value := queries[k][0], queries[k][1], queries[k][2]
            k++
            if right < i {
                continue
            }
            counts[max(left, i)] += value
            counts[right+1] -= value
        }
        sum += counts[i]
    }
    return k
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

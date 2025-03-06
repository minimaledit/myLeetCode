/*
https://leetcode.com/problems/find-missing-and-repeated-values/

You are given a 0-indexed 2D integer matrix grid of size n * n with values in the range [1, n2]. Each integer appears exactly once except a which appears twice and b which is missing.
The task is to find the repeating and missing numbers a and b.
Return a 0-indexed integer array ans of size 2 where ans[0] equals to a and ans[1] equals to b.

Input: grid = [[1,3],[2,2]] Output: [2,4]
Explanation: Number 2 is repeated and number 4 is missing so the answer is [2,4].
*/

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	num := make([]int, n*n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num[grid[i][j]]++
		}
	}
	ans := make([]int, 2)
	for i := 1; i <= n*n; i++ {
		if num[i] == 0 {
			ans[1] = i
		} else {
			if num[i] == 2 {
				ans[0] = i
			}
		}
	}
	return ans
}

/*
https://leetcode.com/problems/maximum-candies-allocated-to-k-children/description/

You are given a 0-indexed integer array candies. Each element in the array denotes a pile of candies of size candies[i].
You can divide each pile into any number of sub piles, but you cannot merge two piles together.
You are also given an integer k. You should allocate piles of candies to k children such that each child gets the same number of candies.
Each child can be allocated candies from only one pile of candies and some piles of candies may go unused.
Return the maximum number of candies each child can get.

Example:
Input: candies = [5,8,6], k = 3  Output: 5
Explanation: We can divide candies[1] into 2 piles of size 5 and 3,
and candies[2] into 2 piles of size 5 and 1. We now have five piles of candies of sizes 5, 5, 3, 5, and 1.
We can allocate the 3 piles of size 5 to 3 children. It can be proven that each child cannot receive more than 5 candies.
*/

func maximumCandies(candies []int, k int64) int {
	left, right, mid := 1, 10000000, 0
	for left <= right {
		mid = (left + right) / 2
		maxPiles := candiesDevide(candies, mid)
		if maxPiles < k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
func candiesDevide(candies []int, num int) (sum int64) {
	for i := 0; i < len(candies); i++ {
		sum += int64(candies[i] / num)
	}
	return
}

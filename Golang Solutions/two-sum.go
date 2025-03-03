/*
https://leetcode.com/problems/two-sum/description/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        difference := target - num
        if idx, ok := numMap[difference]; ok { // if numMap[difference] != nil
            return []int{i, idx}
        }
        numMap[num] = i
    }
    return nil
}

/*
https://leetcode.com/problems/partition-array-according-to-given-pivot/description/

You are given a 0-indexed integer array nums and an integer pivot. Rearrange nums such that the following conditions are satisfied:
- Every element less than pivot appears before every element greater than pivot.
- Every element equal to pivot appears in between the elements less than and greater than pivot.
- The relative order of the elements less than pivot and the elements greater than pivot is maintained.
   - More formally, consider every pi, pj where pi is the new position of the ith element and pj is the new position of the jth element. If i < j and both elements are smaller (or larger) than pivot, then pi < pj.
Return nums after the rearrangement.
*/

func pivotArray1(nums []int, pivot int) []int {
    count := len(nums)
    var left, mid, right []int
    for i:=0; i<count; i++ {
        if(nums[i]<pivot){
            left = append(left, nums[i])
        }else if (nums[i]>pivot){
            right = append(right, nums[i])
        }else {
            mid = append(mid, nums[i])
        }
    }
    return append(append(left, mid...), right...)
}
//with "pointers"
func pivotArray2(nums []int, pivot int) []int {
    var leftBorder int
    rightBorder := len(nums)
    ans := make([]int, rightBorder)
    rightBorder-- //index
    j:=rightBorder
    for i:=0;i<len(nums);i++{
        if(nums[i] < pivot){
            ans[leftBorder] = nums[i]
            leftBorder++
        }
        if(nums[j] > pivot){
            ans[rightBorder] = nums[j]
            rightBorder--
        }
        j--
    }
    for leftBorder != rightBorder + 1{
        ans[leftBorder] = pivot
        leftBorder++
    }
    return ans
}

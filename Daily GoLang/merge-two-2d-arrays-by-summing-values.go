/*
https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/description

You are given two 2D integer arrays nums1 and nums2.

- nums1[i] = [idi, vali] indicate that the number with the id idi has a value equal to vali.
- nums2[i] = [idi, vali] indicate that the number with the id idi has a value equal to vali.
Each array contains unique ids and is sorted in ascending order by id.

Merge the two arrays into one array that is sorted in ascending order by id, respecting the following conditions:

- Only ids that appear in at least one of the two arrays should be included in the resulting array.
- Each id should be included only once and its value should be the sum of the values of this id in the two arrays. If the id does not exist in one of the two arrays, then assume its value in that array to be 0.
Return the resulting array. The returned array must be sorted in ascending order by id.
*/
// First attempt O(N1 + N2 + 1000), simpler solution
func mergeArrays1(nums1 [][]int, nums2 [][]int) [][]int {
	numCounter := make([]int, 1001)
	for i := 0; i < len(nums1); i++ {
		numCounter[nums1[i][0]] += nums1[i][1] // numCounter[id] += value
	}
	for i := 0; i < len(nums2); i++ {
		numCounter[nums2[i][0]] += nums2[i][1] // numCounter[id] += value
	}
	var mergedArray [][]int
	for i := 1; i <= 1000; i++ {
		if numCounter[i] != 0 {
			mergedArray = append(mergedArray, []int{i, numCounter[i]})
		}
	}
	return mergedArray
}
// Second attempt O(N1 + N2)
func mergeArrays2(nums1 [][]int, nums2 [][]int) [][]int {
	len1 := len(nums1)
	len2 := len(nums2)
	i := 0
	j := 0
	var mergedArray [][]int
	for i < len1 || j < len2 {
		if j >= len2 {
			mergedArray = append(mergedArray, []int{nums1[i][0], nums1[i][1]})
			i++
		} else if i >= len1 {
			mergedArray = append(mergedArray, []int{nums2[j][0], nums2[j][1]})
			j++
		} else {
			id1, val1 := nums1[i][0], nums1[i][1]
			id2, val2 := nums2[j][0], nums2[j][1]
			if id1 < id2 {
				mergedArray = append(mergedArray, []int{id1, val1})
				i++
			} else if id1 > id2 {
				mergedArray = append(mergedArray, []int{id2, val2})
				j++
			} else {
				mergedArray = append(mergedArray, []int{id1, val1 + val2})
				i++
				j++
			}
		}
	}
	return mergedArray
}

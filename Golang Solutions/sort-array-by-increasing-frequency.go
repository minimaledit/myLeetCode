/*
https://leetcode.com/problems/sort-array-by-increasing-frequency/description/

Given an array of integers nums, sort the array in increasing order based on the frequency of the values. If multiple values have the same frequency, sort them in decreasing order.
Return the sorted array.

Example 1:
Input: nums = [1,1,2,2,2,3]  Output: [3,1,1,2,2,2]
Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.
*/

type frequency struct {
    value    int
    count    int
}

func frequencySort(nums []int) []int {
    frequencyMap := make(map[int]int)

    for _, num := range nums {
        frequencyMap[num]++
    }

    freqSlice := make([]frequency, 0, len(frequencyMap))
    for value, count := range frequencyMap {
        freqSlice = append(freqSlice, frequency{value: value, count: count})
    }

    sort.Slice(freqSlice, func(i, j int) bool {
        if freqSlice[i].count == freqSlice[j].count {
            return freqSlice[i].value > freqSlice[j].value
        }
        return freqSlice[i].count < freqSlice[j].count
    })

    result := []int{}
    for _, f := range freqSlice {
        for i := 0; i < f.count; i++ {
            result = append(result, f.value)
        }
    }

    return result
}

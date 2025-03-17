/*
https://leetcode.com/problems/sort-characters-by-frequency/description/

Given a string s, sort it in decreasing order based on the frequency of the characters. The frequency of a character is the number of times it appears in the string.
Return the sorted string. If there are multiple answers, return any of them.

Example 1:
Input: s = "tree" Output: "eert"
Explanation: 'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
*/

type frequency struct {
    value    rune
    count    int
}

func frequencySort(s string) string {
    frequencyMap := make(map[rune]int)

    for _, char := range s {
        frequencyMap[char]++
    }

    freqSlice := make([]frequency, 0, len(frequencyMap))
    for value, count := range frequencyMap {
        freqSlice = append(freqSlice, frequency{value: value, count: count})
    }

    sort.Slice(freqSlice, func(i, j int) bool {
        if freqSlice[i].count == freqSlice[j].count {
            return freqSlice[i].value < freqSlice[j].value
        }
        return freqSlice[i].count > freqSlice[j].count
    })

    result := []rune{}
    for _, f := range freqSlice {
        for i := 0; i < f.count; i++ {
            result = append(result, f.value)
        }
    }

    return string(result)
}

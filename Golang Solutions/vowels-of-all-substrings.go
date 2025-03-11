/*
https://leetcode.com/problems/vowels-of-all-substrings/description/

Given a string word, return the sum of the number of vowels ('a', 'e', 'i', 'o', and 'u') in every substring of word.
A substring is a contiguous (non-empty) sequence of characters within a string.
Note: Due to the large constraints, the answer may not fit in a signed 32-bit integer. Please be careful during the calculations.

Input: word = "aba"  Output: 6
Explanation: 
All possible substrings are: "a", "ab", "aba", "b", "ba", and "a".
- "b" has 0 vowels in it
- "a", "ab", "ba", and "a" have 1 vowel each
- "aba" has 2 vowels in it
Hence, the total sum of vowels = 0 + 1 + 1 + 1 + 1 + 2 = 6. 
*/

func countVowels(word string) int64 {
	var totalSum int64
	n := len(word)
	for i, char := range word {
		if isVovel(rune(char)) {
			// Вычисляем количество подстрок, которые включают эту гласную
			totalSum += int64((i+1)*(n-i))
		}
	} 
	return totalSum
}
func isVovel(char rune) bool {
    switch char {
    case 'a', 'e', 'i', 'o', 'u':
        return true
    default:
        return false
    }
}

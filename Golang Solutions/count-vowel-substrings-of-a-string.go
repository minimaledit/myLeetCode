/*
https://leetcode.com/problems/count-vowel-substrings-of-a-string/description/

A substring is a contiguous (non-empty) sequence of characters within a string.
A vowel substring is a substring that only consists of vowels ('a', 'e', 'i', 'o', and 'u') and has all five vowels present in it.
Given a string word, return the number of vowel substrings in word.

Input: word = "cuaieuouac"  Output: 7
Explanation: The vowel substrings of word are as follows:
- "uaieuo"
- "uaieuou"
- "uaieuoua"
- "aieuo"
- "aieuou"
- "aieuoua"
- "ieuoua"
*/

func countVowelSubstrings(word string) int {
	n := len(word)
	count := 0

	for i := 0; i < n; i++ {
		vowelCount := make(map[byte]int)
		for j := i; j < n; j++ {
			char := word[j]
			if isVowel(char) {
				vowelCount[char]++
				if len(vowelCount) == 5 {
					count++
				}
			} else {
				break
			}
		}
	}

	return count
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

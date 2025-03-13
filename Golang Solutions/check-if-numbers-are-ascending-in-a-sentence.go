/*
https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/description/

A sentence is a list of tokens separated by a single space with no leading or trailing spaces. 
Every token is either a positive number consisting of digits 0-9 with no leading zeros, or a word consisting of lowercase English letters.

For example, "a puppy has 2 eyes 4 legs" is a sentence with seven tokens: "2" and "4" are numbers and the other tokens such as "puppy" are words.

Given a string s representing a sentence, you need to check if all the numbers in s are strictly increasing from left to right 
(i.e., other than the last number, each number is strictly smaller than the number on its right in s).
Return true if so, or false otherwise.

Input: s = "1 box has 3 blue 4 red 6 green and 12 yellow marbles"  Output: true
Explanation: The numbers in s are: 1, 3, 4, 6, 12.
They are strictly increasing from left to right: 1 < 3 < 4 < 6 < 12.
*/
func areNumbersAscending(s string) bool {
    s=s+"z"
    arr := make([]int, 0)
	currentNumber := 0
    for i:=0;i<len(s);i++{
        char := s[i]
        if char >= '0' && char <= '9'{
            currentNumber = currentNumber * 10 + int(char) - 48
        }else if currentNumber != 0{
            arr = append(arr, currentNumber)
            currentNumber = 0
        }
    }

    for i:=1;i<len(arr);i++{
        if arr[i-1]>=arr[i]{
            return false
        }
    }
    return true
}

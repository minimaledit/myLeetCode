/*
https://leetcode.com/problems/palindrome-number/description/

Given an integer x, return true if x is a palindrome, and false otherwise.
*/

func isPalindrome(x int) bool {
    //xStr := fmt.Sprintf("%v", x)
    xStr := strconv.Itoa(x)
    for i:=0;i<len(xStr);i++{
        if xStr[i] != xStr[len(xStr) - i - 1]{
            return false
        }
    }

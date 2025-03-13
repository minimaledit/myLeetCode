/*
https://leetcode.com/problems/reverse-integer/description/

Given a signed 32-bit integer x, return x with its digits reversed. 
If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123  Output: 321

Example 2:
Input: x = -123  Output: -321
*/

func reverse(x int) int {
	negative := 1
	if x < 0 {
		x *= -1
		negative = -1
	}
	xStr := strconv.Itoa(x)
	runes := []rune(xStr)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	xReversed, _ := strconv.Atoi(string(runes))
    if xReversed >= -2147483648 && xReversed < 2147483648{
	    return negative * xReversed
    }else{
        return 0
    }
}

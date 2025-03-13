/*
https://leetcode.com/problems/string-to-integer-atoi/description/

Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
The algorithm for myAtoi(string s) is as follows:
- Whitespace: Ignore any leading whitespace (" ").
- Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
- Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
- Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.

Return the integer as the final result.

Example 1:
Input: s = " -042"  Output: -42
Example 2: 
Input: s = "1337c0d3"  Output: 1337
*/
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }
	isPositive := 1
	num := 0
	if s[0] == '-' {
		isPositive = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= 48 && char <= 57 {
			num = num*10 + (int(char) - 48)
		} else {
			break
		}
		if num*isPositive > 2147483647 {
			return 2147483647
		}
		if num*isPositive < -2147483648 {
			return -2147483648
		}
	}
	return isPositive * num
}

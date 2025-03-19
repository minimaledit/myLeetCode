/*
https://leetcode.com/problems/reverse-bits/description/

Reverse bits of a given 32 bits unsigned integer.

Example:
Input: n = 00000010100101000001111010011100  Output: 964176192 (00111001011110000010100101000000)
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.
*/

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans = ans << 1   // shift ans to the left
		ans += (num & 1) // add the last bit to ans
		num = num >> 1   // shift num to the right
	}
	return ans
}

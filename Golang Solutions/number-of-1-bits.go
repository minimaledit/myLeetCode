/*
https://leetcode.com/problems/number-of-1-bits/description/

Given a positive integer n, write a function that returns the number of set bits in its binary representation (also known as the Hamming weight).
*/
func hammingWeight(n int) int {
	//return bits.OnesCount(uint(n))
	var count int
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}

/*
https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/description/

Given two integers left and right, return the count of numbers in the inclusive range [left, right] having a prime number of set bits in their binary representation.
Recall that the number of set bits an integer has is the number of 1's present when written in binary.

For example, 21 written in binary is 10101, which has 3 set bits.
*/

func countPrimeSetBits1(left int, right int) int {
    ans := 0
    for i := left; i <= right; i++ {
        count := bits.OnesCount(uint(i))
        switch count {
        case 2, 3, 5, 7, 11, 13, 17, 19:
            ans++
        }
    }
    return ans
}

// without included function bits.OnesCount
func toBinaryCountOnes(num int) bool {
	var count int
	for num > 0 {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	if count == 2 || count == 3 || count == 5 || count == 7 || count == 11 || count == 13 || count == 17 || count == 19 {
		return true
	} else {
		return false
	}
}
func countPrimeSetBits(left int, right int) int {
	var ans int
	for i := left; i <= right; i++ {
		if toBinaryCountOnes(i) {
			ans++
		}
	}
	return ans
}

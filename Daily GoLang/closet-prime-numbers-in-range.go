/*
https://leetcode.com/problems/closest-prime-numbers-in-range/description/

Given two positive integers left and right, find the two integers num1 and num2 such that:
- left <= num1 < num2 <= right .
- Both num1 and num2 are prime numbers.
- num2 - num1 is the minimum amongst all other pairs satisfying the above conditions.

Return the positive integer array ans = [num1, num2]. If there are multiple pairs satisfying these conditions,
return the one with the smallest num1 value. If no such numbers exist, return [-1, -1].
*/

func closestPrimes(left int, right int) []int {
	if right < 2 {
		return []int{}
	}
	sieve := make([]bool, right+1)
	for p := 2; p*p <= right; p++ {
		if !sieve[p] {
			for i := p * p; i <= right; i += p {
				sieve[i] = true
			}
		}
	}
	var primes []int
	for p := left; p <= right; p++ {
		if !sieve[p] && p >= 2 {
			primes = append(primes, p)
		}
	}
	first, second, minDist := -1, -1, 100000
	for i := 1; i < len(primes); i++ {
		if primes[i]-primes[i-1] < minDist {
			minDist = primes[i] - primes[i-1]
			first = primes[i-1]
			second = primes[i]
		}
	}
	return []int{first, second}
}

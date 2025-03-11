/*
https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/description/

Given a string s consisting only of characters a, b and c.
Return the number of substrings containing at least one occurrence of all these characters a, b and c.

Input: s = "abcabc"  Output: 10
Explanation: The substrings containing at least one occurrence of the characters a, b and c are
"abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again). 
*/
// First Solution: O(N*Log(N))
func numberOfSubstrings(s string) int {
	count, appearA, appearB, appearC := 0, make([]int, 0), make([]int, 0), make([]int, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			appearA = append(appearA, i)
		case 'b':
			appearB = append(appearB, i)
		case 'c':
			appearC = append(appearC, i)
		}
	}
	for i := 0; i < len(s); i++ {
		idxA := lowerBound(appearA, i)
		idxB := lowerBound(appearB, i)
		idxC := lowerBound(appearC, i)
		if idxA == len(appearA) || idxB == len(appearB) || idxC == len(appearC) {
			break
		}
		maxIdx := max(appearA[idxA], appearB[idxB], appearC[idxC])
		
		count += len(s) - maxIdx
	}
	return count
}
func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}

}
// Find Equal or Bigger than target in array
func lowerBound(arr []int, target int) int {
	left, right, mid := 0, len(arr), 0
	for left < right {
		mid = (left + right)/2
		if arr[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}


// Second solution: O(N)
func numberOfSubstrings(s string) int {
	var lastPos = [3]int{-1, -1, -1} 
	var count int
	for i:=0;i<len(s);i++{
		switch s[i]{
		case 'a':
			lastPos[0] = i
		case 'b':
			lastPos[1] = i
		case 'c':
			lastPos[2] = i
		}
		count+=1 + min(lastPos[0], lastPos[1], lastPos[2])
	}
	return count
}
func min(a, b, c int) int {
	if a<=b && a<=c{
		return a
	}else if b<=a && b<=c{
		return b
	}else{
		return c
	}
}



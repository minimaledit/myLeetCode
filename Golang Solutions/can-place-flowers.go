/*
https://leetcode.com/problems/can-place-flowers/description/

You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, 
return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

Input: flowerbed = [1,0,0,0,1], n = 1  Output: true

Input: flowerbed = [1,0,0,0,1], n = 2  Output: false

- 1 <= flowerbed.length <= 2 * 10^4
- flowerbed[i] is 0 or 1.
- There are no two adjacent flowers in flowerbed.
- 0 <= n <= flowerbed.length
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	// Adding zeros, to take into account,
	// that you can plant a flower at the beginning and at the end
	newArr := make([]int, len(flowerbed)+2)
	newArr[0] = 0
	copy(newArr[1:], flowerbed)
	newArr[len(newArr)-1] = 0
	flowerbed = newArr

	i := 0
	ans := 0
	for i < len(flowerbed)-2 {
		if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
			i++
			ans++
		}
		i++
	}
	return (ans >= n)
}

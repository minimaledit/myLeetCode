/*
https://leetcode.com/problems/alternating-groups-ii/description/

There is a circle of red and blue tiles. You are given an array of integers colors and an integer k. The color of tile i is represented by colors[i]:
- colors[i] == 0 means that tile i is red.
- colors[i] == 1 means that tile i is blue.

An alternating group is every k contiguous tiles in the circle with alternating colors (each tile in the group except the first and last one has a different color from its left and right tiles).
Return the number of alternating groups.
Note that since colors represents a circle, the first and the last tiles are considered to be next to each other.

3 <= colors.length <= 10^5
0 <= colors[i] <= 1
3 <= k <= colors.length
*/

func numberOfAlternatingGroups(colors []int, k int) int {
    left, right, ans := 0, 0, 0
    n := len(colors)
    for left < n {
        right++ 
        if colors[right%n] == colors[(right - 1)%n] {
            // If they are the same, the alternating sequence breaks
            // Move the left pointer to the current position to start a new window
            left = right
        }else if right - left + 1 == k {
            // If the window size equals the required group size, it's a valid group
            ans++
            left++
        } 
    }
    return ans
}

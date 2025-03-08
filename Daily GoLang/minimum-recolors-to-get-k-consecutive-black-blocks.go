/*
https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/

You are given a 0-indexed string blocks of length n, where blocks[i] is either 'W' or 'B', representing the color of the ith block. The characters 'W' and 'B' denote the colors white and black, respectively.
You are also given an integer k, which is the desired number of consecutive black blocks.
In one operation, you can recolor a white block such that it becomes a black block.
Return the minimum number of operations needed such that there is at least one occurrence of k consecutive black blocks.

Input: blocks = "WBBWWBBWBW", k = 7    Output: 3
Explanation:
One way to achieve 7 consecutive black blocks is to recolor the 0th, 3rd, and 4th blocks
so that blocks = "BBBBBBBWBW". 
It can be shown that there is no way to achieve 7 consecutive black blocks in less than 3 operations.
*/
// O(N+M)
func minimumRecolors1(blocks string, k int) int {
    var whiteCount int
    minRecolors := 100
    for i:=0;i<=len(blocks)-k;i++{
        for j:=i;j<i+k;j++{
            if blocks[j]=='W'{
                whiteCount++
            }
        }
        if(minRecolors>whiteCount){
            minRecolors = whiteCount
        }
        whiteCount = 0
    }
    return minRecolors
}

//O(N)
func minimumRecolors(blocks string, k int) int {
    var whiteCount, left int
    minRecolors := 100
    for right:=0;right<len(blocks);right++{
        if blocks[right]=='W'{
            whiteCount++
        }
        if right - left == k - 1 {
            if(whiteCount<minRecolors){
                minRecolors = whiteCount
            }
            if blocks[left] == 'W'{
                whiteCount--
            }
            left++
        }
    }
    return minRecolors

/*
https://leetcode.com/problems/pascals-triangle-ii/description/

Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

Example:
Input: rowIndex = 3  Output: [1,3,3,1]
*/

func getRow(rowIndex int) []int {
    pascal := make([][]int,rowIndex+1)
     for i := range pascal {
        pascal[i] = make([]int, i+1)
        pascal[i][0]=1
        pascal[i][i]=1
        for j:=1;j<i;j++{
            pascal[i][j]=pascal[i-1][j-1]+pascal[i-1][j]
        }
    }
    return pascal[rowIndex]
}

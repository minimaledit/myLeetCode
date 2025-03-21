/*
https://leetcode.com/problems/pascals-triangle/description/

Given an integer numRows, return the first numRows of Pascal's triangle.

Example:
Input: numRows = 5  Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
*/

func generate(numRows int) [][]int {
    pascal := make([][]int,numRows)
     for i := range pascal {
        pascal[i] = make([]int, i+1)
        pascal[i][0]=1
        pascal[i][i]=1
        for j:=1;j<i;j++{
            pascal[i][j]=pascal[i-1][j-1]+pascal[i-1][j]
        }
    }
    return pascal
}

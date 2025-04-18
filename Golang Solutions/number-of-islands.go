/*
https://leetcode.com/problems/number-of-islands/description/

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
*/

func numIslands(grid [][]byte) int {
    count := 0
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[i]);j++{
            if(grid[i][j]=='1'){
                dfs(&grid, i, j)
                count++
            }
        }
    }
    return count
}

func dfs(grid *[][]byte, i int, j int){
    if i<0 || j<0 || i>=len(*grid) || j>=len((*grid)[i]) || (*grid)[i][j] != '1'{
        return
    }
    (*grid)[i][j] = '0'

    dfs(grid, i+1,j)
    dfs(grid, i,j+1)
    dfs(grid, i-1,j)
    dfs(grid, i,j-1)
}


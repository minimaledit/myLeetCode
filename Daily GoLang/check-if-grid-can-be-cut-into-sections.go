/*
https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections/description/

You are given an integer n representing the dimensions of an n x n grid, with the origin at the bottom-left corner of the grid.
You are also given a 2D array of coordinates rectangles, where rectangles[i] is in the form [startx, starty, endx, endy],
representing a rectangle on the grid. Each rectangle is defined as follows:
- (startx, starty): The bottom-left corner of the rectangle.
- (endx, endy): The top-right corner of the rectangle.

Note that the rectangles do not overlap. Your task is to determine if it is possible to make either two horizontal or two vertical cuts on the grid such that:
- Each of the three resulting sections formed by the cuts contains at least one rectangle.
- Every rectangle belongs to exactly one section.

Return true if such cuts can be made; otherwise, return false.

Example:
Input: n = 5, rectangles = [[1,0,5,2],[0,2,2,4],[3,2,5,3],[0,4,4,5]]
Output: true
*/

func checkValidCuts(n int, rectangles [][]int) bool {
    var xIntervals, yIntervals [][]int

    for _, rect := range rectangles {
        xIntervals = append(xIntervals, []int{rect[0], rect[2]})
        yIntervals = append(yIntervals, []int{rect[1], rect[3]})
    }

    return check(xIntervals) || check(yIntervals)
}

func check(intervals [][]int) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    sections := 0
    maxEnd := intervals[0][1]

    for _, interval := range intervals {
        if maxEnd <= interval[0] {
            sections++
        }
        maxEnd = max(maxEnd, interval[1])
    }

    return sections >= 2
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

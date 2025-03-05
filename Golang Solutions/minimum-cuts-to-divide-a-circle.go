/*
https://leetcode.com/problems/minimum-cuts-to-divide-a-circle/description/

A valid cut in a circle can be:
-A cut that is represented by a straight line that touches two points on the edge of the circle and passes through its center, or
-A cut that is represented by a straight line that touches one point on the edge of the circle and its center.

Given the integer n, return the minimum number of cuts needed to divide a circle into n equal slices.
*/
func numberOfCuts(n int) int {
    if n==1{
        return 0
    }
    if n%2==0 {
        return n/2
    }
    return n
}

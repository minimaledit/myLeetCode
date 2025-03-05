/*
https://leetcode.com/problems/count-total-number-of-colored-cells/description/

There exists an infinitely large two-dimensional grid of uncolored unit cells. You are given a positive integer n, indicating that you must do the following routine for n minutes:
- At the first minute, color any arbitrary unit cell blue.
- Every minute thereafter, color blue every uncolored cell that touches a blue cell.

Below is a pictorial representation of the state of the grid after minutes 1, 2, and 3.
                  #
          #      ###
    #    ###    #####
          #      ###
                  #
    1     5      13
*/
func coloredCells(n int) int64 {
    n64 := int64(n) 
    //Sn=n(2*a1+(n-1)*d)/2
    //(n-1)*(2*4+((n-1)-1)*4)/2 + 1
    return 2 * n64 * (n64-1) + 1
}

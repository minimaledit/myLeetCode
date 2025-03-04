/*
Given an integer n, return true if it is a power of four. Otherwise, return false.
An integer n is a power of four, if there exists an integer x such that n == 4^x.
*/
func isPowerOfFour(num int) bool {
    for num > 1{
        if num%4 != 0{
            return false
        }
        num/=4
    }
    return (num==1)
}

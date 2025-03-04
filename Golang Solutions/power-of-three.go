/*
Given an integer n, return true if it is a power of three. Otherwise, return false.
An integer n is a power of three, if there exists an integer x such that n == 3^x.
*/
func isPowerOfThree(num int) bool {
    for num > 1{
        if num%3 != 0{
            return false
        }
        num/=3
    }
    return (num==1)
}

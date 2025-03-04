/*
Given an integer n, return true if it is a power of two. Otherwise, return false.
An integer n is a power of two, if there exists an integer x such that n == 2^x.
*/
func isPowerOfTwo1(num int) bool {
    for num > 1{
        if num%2 != 0{
            return false
        }
        num/=2
    }
    return (num==1)
}
//recursion
func isPowerOfTwo2(n int) bool {
    if(n==1){
        return true
    }
    if(n==0){
        return false
    }
    if(n%2==1){
        return false
    }
    return isPowerOfTwo(n/2)
}

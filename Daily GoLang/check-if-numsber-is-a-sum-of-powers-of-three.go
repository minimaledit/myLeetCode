/*
Given an integer n, return true if it is possible to represent n as the sum of distinct powers of three. Otherwise, return false.
An integer y is a power of three if there exists an integer x such that y == 3^x.
Example: Input: n = 12   Output: true
Explanation: 12 = 3^1 + 3^2
*/
// The number can not be represented as a sum of powers of 3 if it's ternary presentation has a 2 in it
func checkPowersOfThree(n int) bool {
    flag := true
    for n!=0 {
        if n % 3 == 2{
            flag = false
            n = 0
        }
        n /= 3
    }
    return flag
}

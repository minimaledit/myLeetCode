/*
https://leetcode.com/problems/add-binary/description/

Given two binary strings a and b, return their sum as a binary string.

Example 1:
Input: a = "11", b = "1" Output: "100"
Example 2:
Input: a = "1010", b = "1011"  Output: "10101"
*/
func addBinary(a string, b string) string {
    lenA, lenB := len(a), len(b)
    if lenA > lenB{
        a,b = b,a
        lenA, lenB = lenB, lenA
    }
    for i:=0;i<lenB-lenA;i++{
        a="0"+a
    }
    a="0"+a
    b="0"+b

    remember := 0
    var s string
    for i:=lenB;i>=0;i--{
        if remember==0{
            if a[i] == '0' && b[i] == '0'{
                s="0"+s
            }else if a[i] == '1' && b[i] == '0'{
                s="1"+s
            }else if a[i] == '0' && b[i] == '1'{
                s="1"+s
            }else if a[i] == '1' && b[i] == '1'{
                s="0"+s
                remember = 1
            }
        }else{
            if a[i] == '0' && b[i] == '0'{
                s="1"+s
                remember = 0
            }else if a[i] == '1' && b[i] == '0'{
                s="0"+s
            }else if a[i] == '0' && b[i] == '1'{
                s="0"+s
            }else if a[i] == '1' && b[i] == '1'{
                s="1"+s
            }
        }
    }
    if s[0]=='0'{
        s=s[1:]
    }
    return s
}

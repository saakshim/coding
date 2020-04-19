/**

https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/start/

PermMissingElem
Find the missing element in a given permutation.

Task description

An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

Your goal is to find that missing element.

Write a function:

func Solution(A []int) int

that, given an array A, returns the value of the missing element.

For example, given array A such that:

  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5
the function should return 4, as it is the missing element.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
the elements of A are all distinct;
each element of array A is an integer within the range [1..(N + 1)].

*/
// Solution 1: Xor of elements with index
// Works better for memory overflow
//
// Logic: i.  X^X = 0
//        ii. 0^X = X
// eg: [3,2,4] = 3^1 ^ 2^2 ^ 4^3
//  Applying (i)&(ii) = 1 ^ 4
//   Final step       = 1 ^ 4 ^ 4(i.e.len(A)+1)
//   Apply (i)&(ii) again, we get = 1

func Solution(A []int) int {
    // write your code in Go 1.4
    xor := 0
    for i, v := range A {
        xor ^= v ^ (i+1)       
    }    
    return xor^(len(A)+1)
}


// Solution 2: Sum of elements in array
// Can run out of memory
 Solution(A []int) int {
    // write your code in Go 1.4
    
    sum, fullSum := 0, 0
    
    for _, v := range A {
        sum += v
    }
    
    for i := 1; i <= len(A)+1; i++ {
        fullSum += i
        
    }
    
    if fullSum > sum {
        return fullSum - sum
    }
    return sum - fullSum
}
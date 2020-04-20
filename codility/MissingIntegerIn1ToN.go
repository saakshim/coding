/**

http://codility.com/demo/take-sample-test/perm_check

PermCheck
Check whether array A is a permutation.
i.e Check if all elements from 1..N are in the array
i.e. Check for missing integer between 1..N in array of N elements

Task description
A non-empty array A consisting of N integers is given.

A permutation is a sequence containing each element from 1 to N once, and only once.

For example, array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
is a permutation, but array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
is not a permutation, because value 2 is missing.

The goal is to check whether array A is a permutation.

Write a function:

func Solution(A []int) int

that, given an array A, returns 1 if array A is a permutation and 0 if it is not.

For example, given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
the function should return 1.

Given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..1,000,000,000].
*/

// Time Complexity: O(N) or O(N * log(N))
//
func Solution(A []int) int {
    // write your code in Go 1.4
    
    size := len(A)
    var x int
    
    for i := 0; i < size; i++ {
        
        x = abs(A[i])
        if x == 0 {
            continue
        }
        if x-1 < size && A[x-1] > 0 {
            A[x-1] = -A[x-1]
        }
    }
    
    for _, v := range A {
        if v > 0 {
            return 0
        }
    }
    return 1
}

func abs(n int) int {
    
    if n < 0 {
        return -n
    }
    return n
}
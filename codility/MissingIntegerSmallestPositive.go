/**

https://codility.com/demo/take-sample-test/missing_integer

https://www.geeksforgeeks.org/find-the-smallest-positive-number-missing-from-an-unsorted-array/

 MissingInteger
Find the smallest positive integer that does not occur in a given sequence.

Task description
This is a demo task.

Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].

*/

// Time Complexity : O(N) or O(N * log(N))
func Solution(A []int) int {
    // write your code in Go 1.4
    
    //split negative and postive
    // all negative on right side and rest are positive
    // shift gives index where positive numbers start    
    A, shift = segregate(A)
    
    var shift int
    var shiftedArr = make([]int, len(A) - shift)
    size := 0
    for i := shift; i < len(A); i++ {
        shiftedArr[size] = A[i]
        size++
    }
    
    var x int
    
    for i:=0; i < size; i++ {
        
        x = abs(shiftedArr[i])
        if x-1 < size && shiftedArr[x-1] > 0 {
            shiftedArr[x-1] = -shiftedArr[x-1]
        }
    }
   
    for i, x := range shiftedArr {
        
        if x > 0 {
            return i+1
        }
    }
    
    return size+1
}

func abs(n int) int {
    
    if n < 0 {
        return -n
    }
    return n
}

func segregate(A []int) ([]int,int) {
    
    j := 0
    tmp := 0
    for i := 0; i < len(A); i++ {
        
        if A[i] <= 0 {
            //swap
            if i != j {
                tmp = A[i]
                A[i] = A[j]
                A[j] = tmp
            }
            j++
        }
    }
    
    return A, j
}
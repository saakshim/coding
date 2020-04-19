/**
Count minimal number of jumps from position X to Y.
https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/start/


A small frog wants to get to the other side of the road. The frog is currently located at position X and wants to get to a position greater than or equal to Y. The small frog always jumps a fixed distance, D.

Count the minimal number of jumps that the small frog must perform to reach its target.

Write a function:

func Solution(X int, Y int, D int) int

that, given three integers X, Y and D, returns the minimal number of jumps from position X to a position equal to or greater than Y.

For example, given:

  X = 10
  Y = 85
  D = 30
the function should return 3, because the frog will be positioned as follows:

after the first jump, at position 10 + 30 = 40
after the second jump, at position 10 + 30 + 30 = 70
after the third jump, at position 10 + 30 + 30 + 30 = 100
Write an efficient algorithm for the following assumptions:

X, Y and D are integers within the range [1..1,000,000,000];
X ≤ Y.

*/

func Solution(X int, Y int, D int) int {
    // write your code in Go 1.4
    
    jumps := 0
    if D == 0 || X >= Y {
        return jumps
    }
        
    
    Y = Y - X
    
    jumps += Y / D


    if Y%D == 0 {
        return jumps
    }
 
    return jumps + 1
}
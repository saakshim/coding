/**
Write a function:

func Solution(S string) int

that, given a string S, returns the index (counting from 0) of a character such that the part of the string to the left of that character is a reversal of the part of the string to its right. The function should return −1 if no such index exists.

Note: reversing an empty string (i.e. a string whose length is zero) gives an empty string.

For example, given a string:

"racecar"

the function should return 3, because the substring to the left of the character "e" at index 3 is "rac", and the one to the right is "car".

Given a string:

"x"

the function should return 0, because both substrings are empty.

Write an efficient algorithm for the following assumptions:

the length of S is within the range [0..2,000,000].

**/


// TimeComplexity:  O(length(S))
func Solution(S string) int {
	// write your code in Go 1.4
	
    if len(S) == 1 {
		return 0
	}

	i := 0
	j := 0
	for j = len(S) - 1; i < j; {


		if S[i] != S[j] {
			break
		} else {
			i++
			j--
		}
	}

	if i == j {
		return i
	}

	return -1
}
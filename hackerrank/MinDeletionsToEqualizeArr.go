/**

Given an array of integers, determine the minimum number of elements to delete to leave only elements of equal value.

Example


Delete the  elements  and  leaving . If both twos plus either the  or the  are deleted, it takes  deletions to leave either  or . The minimum number of deletions is .

Function Description

Complete the equalizeArray function in the editor below.

equalizeArray has the following parameter(s):

int arr[n]: an array of integers
Returns

int: the minimum number of deletions required
Input Format

The first line contains an integer , the number of elements in .
The next line contains  space-separated integers .

Constraints
1 <= n <= 100
1 <= arr[i] <= 100

Sample Input

STDIN       Function
-----       --------
5           arr[] size n = 5
3 3 2 1 3   arr = [3, 3, 2, 1, 3]
Sample Output

2   
Explanation

**/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'equalizeArray' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */


 func equalizeArray(arr []int32) int32 {
    
     var mp = make(map[int32]int32)
     
     var max_freq int32
     max_freq = 0
     
     if len(arr) > 1 {
         max_freq = 1
     }
     for _, x := range arr {
         
         if _, ok := mp[x]; !ok {
             mp[x] = 1
         } else {
             mp[x] = mp[x] + 1
             if max_freq < mp[x] {
                 max_freq = mp[x]
             }
         }
     }
     
     return int32(len(arr)) - max_freq
}

func Abs(x int32) int32 {
    if x < 0 {
        return -x
    }
    return x
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := equalizeArray(arr)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

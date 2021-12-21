/**

An avid hiker keeps meticulous records of their hikes. During the last hike that took exactly  steps, for every step it was noted if it was an uphill, , or a downhill,  step. Hikes always start and end at sea level, and each step up or down represents a  unit change in altitude. We define the following terms:

A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a step down to sea level.
A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level.
Given the sequence of up and down steps during a hike, find and print the number of valleys walked through.

Example

 

The hiker first enters a valley  units deep. Then they climb out and up onto a mountain  units high. Finally, the hiker returns to sea level and ends the hike.

Function Description

Complete the countingValleys function in the editor below.

countingValleys has the following parameter(s):

int steps: the number of steps on the hike
string path: a string describing the path
Returns

int: the number of valleys traversed
Input Format

The first line contains an integer , the number of steps in the hike.
The second line contains a single string , of  characters that describe the path.

Constraints

Sample Input

8
UDDDUDUU
Sample Output

1
Explanation

If we represent _ as sea level, a step up as /, and a step down as \, the hike can be drawn as:

_/\      _
   \    /
    \/\/
The hiker enters and leaves one valley.


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
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */


func countingValleys(steps int32, path string) int32 {
    // Write your code here
    seaLevel := 0
    var valley int32
    valley = 0
    
    // calculate how many times I hit sea level i.e. sealevel = 0
    for _, s := range path {
        
        if s == 'U' {  
            seaLevel += 1    
        } else if s == 'D' {   
            seaLevel -= 1
        } else { 
            return 0
        }
        
        // since we are calculating valleys,
        // how many times do I climb UP & touch sea level
        // gives number of valleys 
        if s == 'U' && seaLevel == 0 {
            valley+=1
        }
    }

    return valley
}

/*

func countingValleys(steps int32, path string) int32 {
    // Write your code here
    foundD := 0
    foundU := 0
    var valley int32
    valley = 0
    for _, s := range path {
        
        if s == 'U' {
            
            if foundD > 0 {
                foundD = foundD - 1
                if foundD == 0 {
                    valley += 1
                }
            } else {
                foundU += 1
            }
        } else if s == 'D' {
            
            if foundU > 0 {
                foundU = foundU - 1
            } else {
                foundD = foundD + 1
            }
        } else {
            
            return 0
        }
        
    }

    return valley
}
*/

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    steps := int32(stepsTemp)

    path := readLine(reader)

    result := countingValleys(steps, path)

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

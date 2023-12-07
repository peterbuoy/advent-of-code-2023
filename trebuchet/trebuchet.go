package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/* Problem Restatement

Go through each line in a text file and find the first and last digit.
Combine the first and last digit (in that order) to form a single two-digit number.
These two-digit numbers are called "calibration values".

The solution is the sum of all calibration values.

// Example
Given a text file and two digit number, we will show the calibration value and the solution:

1abc2			12
pqr3stu8vwx		38
a1b2c3d4e5f		15
treb7uchet		77  # Note that if there is only one number, then it counts as first and last

Solution = 142

Remark: It seems implied that there will always be at least one digit per line.
*/

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file.")
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	/* Strategy
	On a per line basis:
	  - Use left and right pointers that may converge.
	  - If a pointer finds a digit, halt ptr.
	  - If we find a first and a last digit, then create the calibration value and add to sum.
	Edge cases:
	  - Lines with only one digit: left and right ptr will point at same index.
	  - Think about handling single lines of length 1,2, and 3
	*/
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		left, right := 0, len(line)-1
		var first_digit, last_digit string
		var calibrationValue int
		fmt.Println(line)
		// Go's version of for loop where you can have a conditional
		for left <= right {
			// Is there a way to do this without casting?
			if unicode.IsDigit(rune(line[left])) {
				first_digit = string(line[left])
			} else {
				left++
			}
			// Tenacious ternary hate
			if unicode.IsDigit(rune(line[right])) {
				last_digit = string(line[right])
			} else {
				right--
			}
			// Case: found 2 digits in line
			if first_digit != "" && last_digit != "" {
				break
			}
			// Edge Case: only one digit in line
			if left == right {
				// Left ptr grabbed digit first
				if first_digit != "" {
					first_digit = last_digit
				} else { // Right ptr grabbed digit first
					last_digit = first_digit
				}
			}

		}
		calibrationValue, err = strconv.Atoi(first_digit + last_digit)
		fmt.Println(calibrationValue)
		sum += calibrationValue
	}
	fmt.Println("Sum is ", sum)
}

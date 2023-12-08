package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/* Problem Restatement

Go through each line in a text file and find the first and last digit.
This also includes fully spelled out digits like two, which translates to the digit 2
Combine the first and last digit (in that order) to form a single two-digit number.
These two-digit numbers are called "calibration values".

The solution is the sum of all calibration values.

// Example
Given a text file and two digit number, we will show the calibration value and the solution:

1abc2			12
pqr3stu8vwx		38
a1b2c3d4e5f		15
treb7uchet		77  # Note that if there is only one number, then it counts as first and last
one2five9		19	# Note that one counts as the digit 1

Solution = 151

Remark: It seems implied that there will always be at least one digit per line.

Question:
Is the calibration value for 7pqrstsixteen 77?
sixteen isn't a digit...
*/

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file.")
	}
	defer input.Close()

	/* Strategy
	On a per line basis:
	  - Use left and right pointers that may converge.
	  - If a pointer finds a digit, halt ptr.
	  - If we find a first and a last digit, then create the calibration value and add to sum.
	Edge cases:
	  - Lines with only one digit: left and right ptr will point at same index.
	  - Think about handling single lines of length 1,2, and 3
	*/

	spelled_digits := make(map[string]string)
	spelled_digits["one"] = "1"
	spelled_digits["two"] = "2"
	spelled_digits["three"] = "3"
	spelled_digits["four"] = "4"
	spelled_digits["five"] = "5"
	spelled_digits["six"] = "6"
	spelled_digits["seven"] = "7"
	spelled_digits["eight"] = "8"
	spelled_digits["nine"] = "9"

	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {

		line := scanner.Text()
		fmt.Println(line)

		var first_digit, last_digit string
		// left_str, right_str keeps track of spelled digits
		var left_str, right_str string
		var calibrationValue int
		left, right := 0, len(line)-1
		// Damn this is ugly
		for left <= right {
			if first_digit != "" && last_digit != "" {
				break
			}
			if first_digit == "" {
				// Found digit
				if unicode.IsDigit(rune(line[left])) {
					first_digit = string(line[left])
				} else { // Handle letters
					left_str += string(line[left])
					// need to handle values like sixteen -> six
					// check if anything spelled digits is a substring of left_str
					for key, value := range spelled_digits {
						if strings.Contains(left_str, key) {
							first_digit = value
						}
					}
					left++
				}
			}
			if last_digit == "" {
				// Found digit
				if unicode.IsDigit(rune(line[right])) {
					last_digit = string(line[right])
				} else { // Reading from left to right
					right_str = string(line[right]) + right_str
					// Need to handle values like sixteen -> 6
					for key, value := range spelled_digits {
						if strings.Contains(right_str, key) {
							last_digit = value
						}
					}
					right--
				}
			}
		}
		// Edge Cases: only one digit in line

		if last_digit == "" {
			last_digit = first_digit
		}
		if first_digit == "" {
			first_digit = last_digit
		}
		calibrationValue, err = strconv.Atoi(first_digit + last_digit)
		fmt.Println(calibrationValue)
		sum += calibrationValue
	}
	fmt.Println("Sum is ", sum)
}

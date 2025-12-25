package main

import "fmt"

/*
	Given an integer x, return true if x is a palindrome, and false otherwise.

	Example 1:

	Input: x = 121
	Output: true
	Explanation: 121 reads as 121 from left to right and from right to left.
	Example 2:

	Input: x = -121
	Output: false
	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
	Example 3:

	Input: x = 10
	Output: false
	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

	Constraints:

	-2^31 <= x <= 2^31 - 1

	*ADDITIONAL CHALLENGE => Follow up: Could you solve it without converting the integer to a string?*
*/

func isPalindromeA(x int) bool {
	if x < 0 {
		return false
	}
	tmpNum := x
	rev := int(0)
	for tmpNum > 0 {
		d := tmpNum % 10
		rev = rev * 10 + d
		tmpNum = tmpNum / 10
	}
	if rev == x {
		return true
	}
	return false
}

func main() {
	exampleInt := 41314
	res := isPalindromeA(exampleInt)
	fmt.Println(res)
}
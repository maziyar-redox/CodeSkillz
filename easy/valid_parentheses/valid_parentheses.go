package main

import (
	"fmt"
	"strings"
)

/*
	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

	An input string is valid if:

	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
	Every close bracket has a corresponding open bracket of the same type.


	Example 1:

	Input: s = "()"

	Output: true

	Example 2:

	Input: s = "()[]{}"

	Output: true

	Example 3:

	Input: s = "(]"

	Output: false

	Example 4:

	Input: s = "([])"

	Output: true

	Example 5:

	Input: s = "([)]"

	Output: false


	Constraints:

	1 <= s.length <= 104
	s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
    if len(s) % 2 != 0 {
		return false
	}
	hrMap := map[string]int{
		"(": 1,
		")": 2,
		"{": 5,
		"}": 6,
		"[": 9,
		"]": 10,
	}
	prMap := map[string]string{
		"(": ")",
		")": "(",
		"{": "}",
		"}": "{",
		"[": "]",
		"]": "[",
	}
	tmpStr := s
	for i := 0; i < len(tmpStr) - 1; i++ {
		val1, _ := hrMap[string(tmpStr[i])]
		val2, _ := hrMap[string(tmpStr[i + 1])]
		if val1 + 1 == val2 {
			i = i + 1
			continue
		}
		for j := i; j < len(tmpStr); j++ {
			val3, _ := prMap[string(tmpStr[i])]
			lastIndex := strings.LastIndex(tmpStr, val3)
			tmpStr = tmpStr[:indexToRemove] + tmpStr[:lastIndex - 1]
		}
		return false
	}
	return true
}

func main() {
	exampleString := "(){}[]"
	res := isValid(exampleString)
	fmt.Println(res)
}
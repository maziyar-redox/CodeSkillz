package main

import (
	"fmt"
)

/*
	Given an integer n, return true if it is a power of two. Otherwise, return false.

	An integer n is a power of two, if there exists an integer x such that n == 2x.
*/

func generate(numRows int) [][]int {
	// First dimention is for number of rows lenght, seccond is eventually equal to number of length becuase
	// n = r
	arr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]int, i + 1)
		arr[i][0] = 1
		arr[i][len(arr[i]) - 1] = 1
		for j := 1; j < len(arr[i]) - 1; j++ {
			arr[i][j] = arr[i - 1][j - 1] + arr[i - 1][j]
		}
	}
	return arr
}

func main() {
	n := 700
	exmplePascal := generate(n)
	fmt.Println(exmplePascal)
}
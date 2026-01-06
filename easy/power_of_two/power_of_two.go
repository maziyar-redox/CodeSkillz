package main

import "fmt"

/*
	Given an integer n, return true if it is a power of two. Otherwise, return false.
	An integer n is a power of two, if there exists an integer x such that n == 2x.
*/

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 || n == 2 {
		return true
	}
	tmp := 2
	for i := 0; i < n; i++ {
		if tmp == n {
			return true
		}
		if tmp > n {
			break
		}
		tmp *= 2
	}
	return false
}

func main() {
	fmt.Println(isPowerOfTwo(8))
}
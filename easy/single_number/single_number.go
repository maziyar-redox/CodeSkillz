package main

import "fmt"

/*
	Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
	You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

func singleNumber(nums []int) int {
	x := nums[0]
	for i := 0; i < len(nums); i++ {
		nums[i] = x ^ nums[i]
	}
	return x
}

func main() {
	exmp := []int{4,1,2,1,2}
	fun := singleNumber(exmp)
	fmt.Println("var is", fun)
	//fmt.Println(val)
}
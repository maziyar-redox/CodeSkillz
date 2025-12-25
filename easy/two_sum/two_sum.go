package main

import "fmt"

/*
	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.
	You can return the answer in any order.

	Example 1:

	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	Example 2:

	Input: nums = [3,2,4], target = 6
	Output: [1,2]
	Example 3:

	Input: nums = [3,3], target = 6
	Output: [0,1]

	Constraints:

	2 <= nums.length <= 104
	-10^9 <= nums[i] <= 10^9
	-10^9 <= target <= 10^9
	Only one valid answer exists.

	*ADDITIONAL CHALLENGE => Follow-up: Can you come up with an algorithm that is less than O(n^2) time complexity?*
*/

// The function bellow twoSumA time complexity is O(n^2) because it runs two for loop on a
// Slice.

func twoSumA(nums []int, target int) []int {
	for indexI, valI := range nums {
		for indexJ, valJ := range nums {
			if valI + valJ == target && indexI != indexJ {
				return []int{indexI, indexJ}
			}
		}
	}
	return []int{}
}

// The function uses hash table to reduce O(n^2) complexity to O(n) and space complexity increased from O(1) to O(n)
// In this method

func twoSumB(nums []int, target int) []int {
	tmp := make(map[int]int)
	for index, val := range nums {
		complement := target - val
		if j, ok := tmp[complement]; ok {
			return []int{j, index}
		}
		tmp[val] = index
	}
	return []int{}
}

func main() {
	exampleSlice := []int{2, 4, 1, 5}
	target := 9
	isAddUpA := twoSumA(exampleSlice, target)
	fmt.Println(isAddUpA)
	isAddUpB := twoSumB(exampleSlice, target)
	fmt.Println(isAddUpB)
}
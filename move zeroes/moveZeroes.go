package main

import "fmt"

/*
Given an integer array nums, move all 0's to the end of it
while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
*/

func moveZeroes(nums []int) {
	nonZeroIndex := 0

	for i, num := range nums {
		if num != 0 {
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			nonZeroIndex++
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5}
	moveZeroes(arr)
	fmt.Println(arr)
}

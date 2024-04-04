package main

import "fmt"

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

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
*/

// higher time complexity
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// higher memory usage, less time complexity
func _twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		value, ok := hashMap[complement]
		if ok {
			return []int{value, i}
		}
		hashMap[num] = i
	}

	return nil
}

func __twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		value, ok := hashMap[target-num]
		if ok {
			return []int{value, i}
		}
		hashMap[num] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(_twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(__twoSum([]int{2, 7, 11, 15}, 9))
}

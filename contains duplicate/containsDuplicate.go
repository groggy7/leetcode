package main

import "sort"

/*
Given an integer array nums, return true if any value appears at least
twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

//My first solution
/* Probably the worst one :)
func containsDuplicate(nums []int) bool {
    counter := make(map[int]int)

    for _, num := range nums {
        counter[num]++
    }

    for _, value := range counter {
        if value > 1 {
            return true
        }
    }

    return false
}*/

//Second solution, this uses bool instead of int
/*func containsDuplicate(nums []int) bool {
	counter := make(map[int]bool)

	for _, num := range nums {
		if counter[num] {
			return true
		}

		counter[num] = true
	}

	return false
}*/

// Third solution, first we sort the array, then check
// if the indices are the same, if so return true
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

//Fourth solution, hash set
/*func containsDuplicate(nums []int) bool {
    seen := make(map[int]struct{})

    for _, num := range nums {
        if _, exists := seen[num]; exists {
            return true
        }
        seen[num] = struct{}{}
    }
    return false
}*/

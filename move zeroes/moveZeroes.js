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

function moveZeroes(arr) {
    let nonZeroIndex = 0;

    for (let i = 0; i < arr.length; i++) {
        if (arr[i] !== 0) {
            arr[nonZeroIndex] = arr[i];
            nonZeroIndex++;
        }
    }

    for (let i = nonZeroIndex; i < arr.length; i++) {
        arr[i] = 0;
    }

    return arr;
}

console.log(moveZeroes([0,0,1]))

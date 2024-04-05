//high time complexity, average memory usage
function containsDuplicate(nums) {
    nums.sort()

    for(i = 0; i < nums.length; i++) {
        if(nums[i] == nums[i+1]) {
            return true
        }
    }
    return false
}

// lower time complexity, lower memory usage
function _containsDuplicate(nums) {
   let set = new Set()

   for(i = 0; i < nums.length; i++) {
    if(set.has(nums[i])) {
        return true
    }
    set.add(nums[i])
   }
   
}

console.log(containsDuplicate([1, 2, 3, 1]))
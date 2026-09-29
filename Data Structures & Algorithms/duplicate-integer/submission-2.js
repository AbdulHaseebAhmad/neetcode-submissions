class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
          let numbers = new Set(nums)          
          return numbers.size !== nums.length 
    }
}

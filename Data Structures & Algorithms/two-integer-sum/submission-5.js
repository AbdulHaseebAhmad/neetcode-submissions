class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */

    twoSum(nums, target) {
        let i = 0;
        while (i < nums.length - 1) {
        let j = nums.length - 1;
            while (j > i) {
                if (nums[i] + nums[j] == target) {
                    return [i,j]
                }
                j--;
            }
            i++;
        }
        return [];
    }
}

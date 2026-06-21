func hasDuplicate(nums []int) bool {
    for index, number := range nums {
        for j := index + 1; j < len(nums); j++ {
            if number == nums[j] {
                return true
            }
        }
    }
    return false
}
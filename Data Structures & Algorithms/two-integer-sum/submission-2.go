func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	var result []int
	for index,value := range nums {
		remainder := target - value
		sval, exists :=  seen[remainder]
		if exists {
			result = append(result, sval)
			result = append(result, index)
		} else {
			seen[value] = index
		}
	}
	return result
}

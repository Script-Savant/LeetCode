func twoSum(nums []int, target int) []int {
	prev := map[int]int{}

	for i, num := range nums {
		complement := target - num

		index, ok := prev[complement]
        if ok {
            return []int {i, index}
        }
        prev[num] = i
	}

	return []int{}
}
func twoSum(nums []int, target int) []int {
    indexes := make(map[int]int)

    for index, num := range nums {
        comp :=  target - num

        idx, ok := indexes[comp]

        if ok {
            return []int{idx, index}
        } else {
            indexes[num] = index
        }
    }
    return nil
}
func removeElement(nums []int, val int) int {
    fmt.Println(nums)
    i:= 0
    for j := 0; j < len(nums); j++{
        if nums[j] != val{
            nums[i] = nums[j]
            i++
        }
    }

    fmt.Println(nums)

    return i
}
func removeDuplicates(nums []int) int {

    n := len(nums)
  
    if n == 0 {
        return 0
    }

    counter := 0

    for i := 1; i < n; i++{
        if nums[i] != nums[counter] {
            counter++
            nums[counter] = nums[i]
        }
    }

    return counter + 1
}
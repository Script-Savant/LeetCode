func containsDuplicate(nums []int) bool {
    // Alex
    seen := make(map[int] bool)
    
    for _, num := range nums{
        if seen[num]{
            return true
        }
        seen[num] = true
    }
    return false
}
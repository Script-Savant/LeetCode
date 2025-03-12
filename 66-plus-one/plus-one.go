func plusOne(digits []int) []int {
    n := len(digits)
    if n == 0{
        return []int {}
    }

    for i := n-1; i >= 0; i--{
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }

    digits = append([]int {1}, digits...)
    return digits
}
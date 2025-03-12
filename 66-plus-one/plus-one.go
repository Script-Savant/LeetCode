func plusOne(digits []int) []int {
    fmt.Println(digits)
    n := len(digits)
    if n == 0{
        return []int {}
    }

    for i := n-1; i >= 0; i--{
        if digits[i] < 9 {
            digits[i]++
            fmt.Println(digits)
            return digits
        }
        digits[i] = 0
        fmt.Println(digits)
    }

    digits = append([]int {1}, digits...)
    fmt.Println(digits)
    return digits
}
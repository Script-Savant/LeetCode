func isPalindrome(x int) bool {
    if x < 0{
        return false
    }

    if x != 0 && x % 10 == 0{
        return false
    }

    original := x
    reversed := 0

    for original > 0{
        reversed = reversed * 10 + original % 10
        original /= 10
    }

    return x == reversed
}
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }

    start, end := 0, 0
    
    expandAroundCenter := func(left, right int) (int, int) {
        for left >= 0 && right < len(s) && s[left] == s[right] {
            left--
            right++
        }
        return left + 1, right - 1
    }

    for i := 0; i < len(s); i++ {
        // Odd-length palindrome
        l1, r1 := expandAroundCenter(i, i)
        // Even-length palindrome
        l2, r2 := expandAroundCenter(i, i+1)

        if r1-l1 > end-start {
            start, end = l1, r1
        }
        if r2-l2 > end-start {
            start, end = l2, r2
        }
    }

    return s[start : end+1]
}
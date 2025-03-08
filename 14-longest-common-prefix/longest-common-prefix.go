func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    // find the shortest string in the slice
    shortest := strs[0]
    for _, s := range strs {
        if len(s) < len(shortest){
            shortest = s
        }
    }

    // Iterate over the characters of the shortest string
    for i := 0; i < len(shortest); i++{
        char := shortest[i]
        for _, s := range strs {
            if s[i] != char {
                return shortest[:i]
            }
        }
    }

    return shortest
}
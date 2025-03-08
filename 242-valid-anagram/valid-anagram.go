func isAnagram(s string, t string) bool {
    // Alex
    if len(s) != len(t){
        return false
    }

    freq := make(map[rune]int)

    for _, value := range s{
        freq[value] += 1
    }

    for _, value :=range t{
        freq[value] -= 1
    }

    for _, value := range freq{
        if (value) != 0{
            return false
        }
    }

    return true
}
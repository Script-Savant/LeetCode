/*
split the input string by space
initialize the left to 0 and the right to len(words) - 1
Traverse the elements in the words list using the right and left pointers until the pointers meet
Each iteration
    - Swap the element pointed by the right and left pointer
    - increment left
    - decrement right
*/
// import "fmt"
func reverseWords(s string) string {
    // fmt.Println(s)

    s = strings.TrimSpace(s)
    words := strings.Fields(s)

    // fmt.Println(words)

    left := 0
    right := len(words) - 1

    for left < right{
        words[left], words[right] = words[right], words[left]
        left++
        right--
    }

    result := strings.Join(words, " ")

    return result

}
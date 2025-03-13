func fullJustify(words []string, maxWidth int) []string {
    var result []string
    n := len(words)
    i := 0

    for i < n {
        // Start a new line with the current word
        lineLength := len(words[i])
        j := i + 1

        // Add as many words as possible to the current line
        for j < n && lineLength+1+len(words[j]) <= maxWidth {
            lineLength += 1 + len(words[j])
            j++
        }

        // Calculate the number of words in the current line
        numWords := j - i
        totalSpaces := maxWidth - (lineLength - (numWords - 1)) // Total spaces needed

        // Build the justified line
        var line strings.Builder
        if j == n || numWords == 1 {
            // Left-justify for the last line or lines with a single word
            line.WriteString(words[i])
            for k := i + 1; k < j; k++ {
                line.WriteString(" ")
                line.WriteString(words[k])
            }
            // Pad the remaining spaces on the right
            line.WriteString(strings.Repeat(" ", maxWidth-line.Len()))
        } else {
            // Distribute spaces evenly for fully-justified lines
            spacesBetweenWords := totalSpaces / (numWords - 1)
            extraSpaces := totalSpaces % (numWords - 1)

            for k := i; k < j; k++ {
                line.WriteString(words[k])
                if k < j-1 {
                    // Add the base spaces
                    line.WriteString(strings.Repeat(" ", spacesBetweenWords))
                    // Add extra spaces if any
                    if extraSpaces > 0 {
                        line.WriteString(" ")
                        extraSpaces--
                    }
                }
            }
        }

        // Add the line to the result
        result = append(result, line.String())
        i = j
    }

    return result
}
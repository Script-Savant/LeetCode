// 1. initialize candies array with 1 candy for each child
// 2. left to right pass - if a child's rating is higher than their left neighbour, give them one more candy
// 3. right to left pass - if a child's rating is higher than their left neighbour, give them one more candy
// 4. Sum the candies
func candy(ratings []int) int {
    n := len(ratings)
    if n == 0 {
        return 0
    }

    candies := make([]int, n)
    for i := range candies{
        candies[i] = 1
    }

    for i := 1; i < n; i++{
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1] + 1
        }
    }

    for i := n -2; i >= 0; i-- {
        if ratings[i] > ratings[i+1]{
            if candies[i] <= candies[i+1] {
                candies[i] = candies[i+1] + 1
            }
        }
    }

    total := 0
    for _, num := range candies {
        total += num
    }

    return total

}
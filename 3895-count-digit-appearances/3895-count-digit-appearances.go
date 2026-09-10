func countDigitOccurrences(nums []int, digit int) int {
    count := 0
    for _, val := range nums {
        for val > 0 {
            valdig := val % 10
            if valdig == digit {count++}
            val /= 10
        }
    }
    return count
}
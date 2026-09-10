func countDigitOccurrences(nums []int, digit int) int {
    count := 0
    for _, val := range nums {
        count += strings.Count(strconv.Itoa(val), strconv.Itoa(digit))
    }
    return count
}
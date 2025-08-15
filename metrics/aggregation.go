package metrics

// Sum returns the sum of a slice of ints.
func Sum(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

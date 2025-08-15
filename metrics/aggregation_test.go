package metrics

import "testing"

func TestSum(t *testing.T) {
    cases := []struct {
        input []int
        want  int
    }{
        {[]int{}, 0},
        {[]int{1, 2, 3}, 6},
        {[]int{-1, 1}, 0},
    }
    for _, c := range cases {
        got := Sum(c.input)
        if got != c.want {
            t.Fatalf("Sum(%v) = %d, want %d", c.input, got, c.want)
        }
    }
}

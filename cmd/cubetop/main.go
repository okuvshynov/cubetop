package main

import (
    "fmt"
    "github.com/okuvshynov/cubetop/metrics"
)

func main() {
    // Example usage of metrics package
    data := []int{1, 2, 3, 4, 5}
    sum := metrics.Sum(data)
    fmt.Printf("Sum of %v is %d\n", data, sum)
}

package main

import "fmt"

func sum(values []int) int {
	total := 0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	return total
}

func main() {
	result := sum([]int{10, 20, 30})
	fmt.Println("The sum is:", result)
}

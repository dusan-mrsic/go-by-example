package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	sum1 := sum(1, 2, 3)
	fmt.Println(sum1)

	sum2 := sum(1, 2, 3, 4)
	fmt.Println(sum2)

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...))
}

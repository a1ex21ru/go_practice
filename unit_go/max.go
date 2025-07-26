package main

import "fmt"

func main() {
	fmt.Println(Max([]int{1, 2, 45, 7, 87, 74}))
}

func Max(numbers []int) int {
	max := 0

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func MaxIndex(numbers []int) int {
	max := 0
	index := -1

	for i, num := range numbers {
		if num > max {
			index = i
			max = num
		}
	}
	return index
}

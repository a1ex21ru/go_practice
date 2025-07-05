package main

import "fmt"

func sum(n uint) uint {

	if n == 1 {
		return n
	}
	return n + sum(n-1)
}

func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func recursion() {

	fmt.Println(sum(6))

	fmt.Println(factorial(6))
}

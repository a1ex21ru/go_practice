package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	for i := 1; i < 7; i++ {
		go sum(i)
	}

	fmt.Println(runtime.GOMAXPROCS(0))
	time.Sleep(5 * time.Millisecond)
	fmt.Println("The End")
}

func sum(n int) {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	fmt.Println(n, "-", result)
}

package main

import "fmt"

func main() {

	list := make([]int, 4, 5)

	list = append(list, 1)

	list[0] = 5
	//list2[0] = 9

	fmt.Println(list)
	//fmt.Println(list2)

}

func handle(list []int) {
	//list[1] = 123
	res := append(list, 342)
	fmt.Println("func", res)
}

func double(nums []int) []int {
	var res []int = make([]int, len(nums))

	for i, num := range nums {
		res[i] = num * 2
	}
	return res
}

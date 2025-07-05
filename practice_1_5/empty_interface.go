package main

import "fmt"

type Empty interface{}

func print_value(value Empty) {
	fmt.Println(value)
}

type person struct {
	name string
}

type account struct {
	id int
}

func main() {

	tom := person{"Tom"}
	tom_acc := account{125}

	print_value(tom_acc)
	print_value(tom)

	//num := 23

	arr := [][]int{{1, 2}, {3, 4}}

	print_value(len(arr[0]))
}

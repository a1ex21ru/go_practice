package main

import "fmt"

type binaryOp func(int, int) int

func action(n1 int, n2 int, op binaryOp) {

	result := op(n1, n2)
	fmt.Println(result)
}

func add(x int, y int) int {
	z := x + y
	return z
}

func pointer() {
	var x = 4

	var (
		a = 2
		b = 3
		c = 4
	)

	var p *int = &x

	fmt.Println(p, *p)

	p1 := new(int)
	fmt.Println(*p1)

	*p1 = 54
	fmt.Println(*p1)

	var p2 **int = &p
	fmt.Println(p2, **p2)

	var p_nums [4]*int = [...]*int{&a, &b, nil, &c}

	//fmt.Println(p_nums)

	for _, value := range p_nums {
		if value != nil {
			fmt.Print(*value, "\t")
		} else {
			fmt.Print("Nil\t")
		}
	}
	fmt.Print("\n")

	my_Op := add
	action(10, 35, my_Op)

}

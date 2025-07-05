package main

import "fmt"

func outer() func() { // внешняя функция
	var n int = 5     // некоторая переменная - лексическое окружение функции inner
	inner := func() { // вложенная функция
		// действия с переменной n
		n += 1
		fmt.Println(n)
	}
	return inner
}

func multiply(n int) func(int) int {

	return func(m int) int { return n * m }
}

func closure() {

	fn := outer() // fn = inner, так как функция outer возвращает функцию inner
	// вызываем внутреннюю функцию inner
	fn() // 6

	func_res := multiply(44)
	res1 := func_res(3)
	fmt.Println(res1)
}

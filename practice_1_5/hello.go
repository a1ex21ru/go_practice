package main

import (
	"fmt"
	"strings"
)

func my_hello() {

	str1 := "Привет мир"

	// выводим в символьном виде
	for _, value := range str1 {
		fmt.Printf("%c ", value)
	}
	fmt.Print("\n")

	// выводим числовые значения байтов строки
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%d ", str1[i])
	}
	fmt.Print("\n")

	// выводим числовые значения байтов строки в символьном виде
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i])
	}
	fmt.Print("\n")
}

func rune_slice() {
	str2 := "Привет мир"

	rune_slice := []rune(str2) // преобразуем str2 в срез рун
	for i := 0; i < len(rune_slice); i++ {
		fmt.Printf("%c     ", rune_slice[i])
	}
	fmt.Println()
	for i := 0; i < len(rune_slice); i++ {
		fmt.Printf("%d  ", rune_slice[i])
	}
	fmt.Print("\n")
}

func heloo() {
	fmt.Println("Hello world")
}

func add_hel(numbers ...int) {
	fmt.Println("function")
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("sum = ", sum)
}

func run_str() {

	str := "Hello"

	fmt.Println(str) // Hello

	rune_str := []rune(str) // преобразуем строку в руны
	rune_str[1] = 'a'       // изменяем руны

	str = string(rune_str) // преобразуем обратно руны в строку

	fmt.Println(str) // Hallo
}

var f int = 52

func main() {

	//// Переменные
	var f1, f2, f3 float32
	f1 = 2.5
	f2 = 3.2
	f3 = 5.5

	one, two := 1, 2

	var (
		// 	zetta = 5
		alpha = 65 - (-2)
	)

	f1 -= f2
	f3 += float32(one)
	two *= int(f1)

	const pi float64 = 3.1415

	user_name := "Эцио Аудиторе"

	fmt.Printf("Name's type: %T\n", user_name)

	fmt.Printf("alpha(%%c): %c", alpha)
	fmt.Println('\n')

	numbers := []int{1, 2, 0, 0, 5, 3}
	fmt.Println(numbers[0])

	add_hel(numbers...)

	numbers1 := [...]int{3, 4, 5, 6, 5, 7, 88}

	fmt.Println("Numbers length:", len(numbers1)) // Numbers length: 4

	if len(numbers1) < len(numbers) {
		fmt.Print("length numbers win")
	} else {
		fmt.Print("length numbers1 win\n")
	}

	switch pi {
	case 2:
		fmt.Println("a = 2")
	case 1:
		fmt.Println("a = 1")
	default:
		fmt.Println("not defined")
	}

	for i := 1; i < 10; i++ {
		fmt.Println((i + 1) * i)
	}

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	str := "hello"
	for index, value := range str {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}

	// метка для перехода к внешнему циклу
	// OuterLoop:
	// 	for i := 1; i <= 3; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			fmt.Printf("i = %d, j = %d\n", i, j)
	// 			if i == 2 && j == 2 {
	// 				fmt.Println("Выход из внешнего цикла...")
	// 				break OuterLoop // выходим из внешнего цикла
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("Цикл завершен...")

	var f func(int, int) int = func(x, y int) int { return x + y }

	fmt.Println(f(2, 6))

	fmt.Println()

	run_str()

	var str_lang string = "golang c++ python csharp"

	str_lang = strings.ToUpper(str_lang[2:])

	fmt.Println(str_lang)

	fmt.Println(strings.Trim(str_lang, "LARP"))
}

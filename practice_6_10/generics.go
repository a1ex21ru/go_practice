package main

import (
	"fmt"

	. "golang.org/x/exp/constraints"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person // Встроенная структура
	Salary float64
}

func generics() {

	// Использование:
	e := Employee{
		Person: Person{Name: "Alice", Age: 30},
		Salary: 100000,
	}
	e.print()

	//list := []int{1, 2, 3, 4}
	//listAge := []int{15, 24, 75, 19}
	////fmt.Println(IsContains(4, list))
	////fmt.Println(IsContains(44, list))
	//
	//list2 := []string{"apple", "banana", "malone", "orange"}
	//fmt.Println(IsContains("apple", list2))
	//
	////sum := func(x, y int) int { return x + y }
	//multy := func(x, y int) int { return x * y }
	//
	//res := Reduce(list, multy, 1)
	//fmt.Println(res)
	//
	//x2 := func(x int) int { return x * 2 }
	//fmt.Println(Map(list, x2))
	//
	//filterFunc := func(x int) bool {
	//	if x < 18 {
	//		return false
	//	}
	//	return true
	//}
	//fmt.Println(Filter(listAge, filterFunc))
	// m := make(map[int]int)
	// m[1] = 5
	// m[0] = 4
	// m[2] = 1
	// m[142] = 66
	// fmt.Println(m)
}

func (empl Employee) print() {
	fmt.Println(empl.Name, empl.Age, empl.Salary)
}

func AppendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func CopySlice(sl, cp []int) {
	copy(sl, cp)
}

func MutateSlice(sl []int, i, val int) {
	sl[i] = val

func main() {
	list := []int{1, 2, 3, 4}
	listAge := []int{15, 24, 75, 19}
	//fmt.Println(IsContains(4, list))
	//fmt.Println(IsContains(44, list))

	list2 := []string{"apple", "banana", "malone", "orange"}
	fmt.Println(IsContains("apple", list2))

	//sum := func(x, y int) int { return x + y }
	multy := func(x, y int) int { return x * y }

	res := Reduce(list, multy, 1)
	fmt.Println(res)

	x2 := func(x int) int { return x * 2 }
	fmt.Println(Map(list, x2))

	filterFunc := func(x int) bool {
		if x < 18 {
			return false
		}
		return true
	}
	fmt.Println(Filter(listAge, filterFunc))

func IsContains[T comparable](n T, list []T) bool {
	for _, v := range list {
		if v == n {
			return true
		}
	}
	return false
}

func Max[T Ordered](x, y T) T {

	if x > y {
		return x
	}
	return y
}

type Number interface {
	Integer | Float
}

func Sum[T Number](list []T) T {
	var sum T
	for _, v := range list {
		sum += v
	}
	return sum
}

func Reduce[T any](list []T, accumulator func(T, T) T, init T) T {
	for _, v := range list {
		init = accumulator(init, v)
	}

	return init
}

func Filter[T any](list []T, f func(T) bool) []T {
	var result []T
	for _, v := range list {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T any](list []T, f func(T) T) []T {
	var result []T
	for _, v := range list {
		result = append(result, f(v))
	}
	return result
}

package main

import "fmt"

type Reader interface {
	read()
}

type File struct {
	text string
}

// func (f File) read() {
// 	fmt.Println(f.text)
// }

func (f *File) read() {
	fmt.Println(f.text)
}

func read_data(data Reader) {
	data.read()
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case float64:
		fmt.Println("float64:", v)
	case File:
		fmt.Println("File: ", v)
	default:
		fmt.Println("unknown type")
	}
}

func main() {

	file := File{"Hello world"}

	//read_data(file)

	var p_file *File = &file

	read_data(p_file)

	checkType(file)
}

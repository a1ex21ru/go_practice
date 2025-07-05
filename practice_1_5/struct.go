package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type person struct {
	name    string
	surname string
	age     int
}

type account struct {
	login    string
	password string

	person_info person
}

// Методы структур

func (p person) print() {
	fmt.Println("Name: ", p.name, "\nAge: ", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "eat's", meal)
}

func (p *person) change_age(new_age int) {
	p.age = new_age
}

func helo(pers person) {
	fmt.Println("Hello ", pers.name, "\nYour birth year: ", 2025-pers.age)
}

func system(acc account) {
	fmt.Println("You are welcome! Mister ", acc.person_info.surname, "\n",
		"your login: ", acc.login, "\n", "your password: ", acc.password)
	helo(acc.person_info)
}

func main() { //struct

	heloo()

	tom := person{name: "Holland", age: 24}

	helo(tom)

	hidl := tom
	hidl.age = 40

	helo(hidl)

	oliver := account{
		login:    "greenarrow.com",
		password: "152347934",
		person_info: person{
			name:    "Oliver",
			surname: "Queen",
			age:     30,
		},
	}

	oliver.password = "1548792637a"

	//system(oliver)

	fmt.Println()
	oliver.person_info.print()

	//oliver.person_info.eat("fuagra")

	oliver.person_info.change_age(40)

	oliver.person_info.print()

}

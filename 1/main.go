package main

import (
	"errors"
	"fmt"
	//	. "golang-ninja/basic/shape"
)

const winePrice = 100

func main() {
	// user := User{"Vasya", 23, "male", 78, 180}
	// users := []User{user}
	// users = append(users, User{"fot", 55, "male", 44, 100})
	//user.printff()

	change, err := buy_wine(177, 101)
	if err != nil {
		fmt.Println("failed: ", err)
	} else {
		fmt.Println("your change: ", change)
	}

}

func buy_wine(age, money int) (int, error) {
	if age >= 18 && money >= winePrice {
		return money - winePrice, nil
	} else if age < 18 {
		return money, errors.New("sorry, you are little")
	} else {
		return money, errors.New("don't enougt money")
	}
}

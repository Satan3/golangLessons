package main

import "fmt"

func main() {
	// default value
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer)

	// nil pointer panic
	//fmt.Printf("%T %#v %#v \n", intPointer, intPointer, *intPointer)

	// Получение not-nil указателей
	// variable
	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	// get variable pointer
	var pointerA *int64 = &a
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA)

	// get pointer via new keyword
	var newPointer = new(float32)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 3
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)

	// pointers usage
	// side effects
	num := 3
	square(num)
	fmt.Println(num)

	squarePointer(&num)
	fmt.Println(num)

	// empty value flag
	var wallet1 *int
	fmt.Println(hasWallet(wallet1))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 100
	fmt.Println(hasWallet(&wallet3))
}

func square(num int) {
	num *= num
}

func squarePointer(num *int) {
	*num = *num * *num
}

func hasWallet(money *int) bool {
	return money != nil
}

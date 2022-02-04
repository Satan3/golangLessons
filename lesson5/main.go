package main

import "fmt"

func main() {
	var firstVariable bool
	fmt.Println(firstVariable)

	secondVariable := true
	fmt.Println(secondVariable)

	age := 15

	// Basic if
	if (age < 18) == true { // можно короче: if age < 18 {
		fmt.Println("You are too young (full)")
	}

	//Short syntax
	if isChild := isChildren(age); isChild == true { // можно короче: if isChild := isChildren(age); isChild {
		fmt.Println("You are too young (short)")
		fmt.Println(isChild)
	}
	//fmt.Println(isChild)

	// If ... else
	age = 20
	if age < 18 {
		fmt.Println("you are too young")
	} else {
		fmt.Println("You are an adult")
	}

	// &&
	if age >= 7 && age <= 18 {
		fmt.Println("You are in school")
	}

	// ||
	if age == 14 || age == 20 || age == 40 {
		fmt.Println("You have to get a new passport")
	}

	// !
	if !isChildren(age) {
		fmt.Println("You are an adult")
	}
}

func isChildren(age int) bool {
	return age < 18
}

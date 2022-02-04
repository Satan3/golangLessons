package main

import "fmt"

/*
func funcName (params...) (returned values) {
	function`s body
}
*/

func main() {
	first, second := 1, 2
	Greet()

	PersonGreet("Vasya")

	FioGreet("John", "Smith")

	summa := Sum(first, second)
	fmt.Println(summa)

	summa, multiply := SumAndMultiply(first, second)
	fmt.Println(summa, multiply)

	_, multiply64 := namedSumAndMultiply(first, second)
	fmt.Println(multiply64)

	/* Func as values */
	var multiplier func(x, y int) int

	multiplier = func(x, y int) int { return x * y }
	fmt.Println(multiplier(first, second))

	divider := func(x, y int) int { return x / y }
	fmt.Println(divider(second, first))

	// Pass func as an argument
	sumFunc := func(x, y int) int {
		return x + y
	}

	subtractFunc := func(x, y int) int {
		return x - y
	}

	fmt.Println(calculate(first, second, sumFunc))
	fmt.Println(calculate(first, second, subtractFunc))

	// Return new function

	divideBy2 := createDivider(2)
	divideBy10 := createDivider(10)

	fmt.Println(divideBy2(100))
	fmt.Println(divideBy10(100))

	// Closures

	dollar := 10

	getDollarValue := func() int {
		return dollar
	}

	fmt.Println(getDollarValue())
	dollar = 70

	fmt.Println(getDollarValue())
}

// Simple hello guys function

func Greet() {
	fmt.Println("Hello guys")
}

// Func with arguments

func PersonGreet(name string) {
	fmt.Printf("Zdarova %s\n", name)
}

func FioGreet(name, surname string) {
	fmt.Printf("Hi %s %s\n", name, surname)
}

// Func with return

func Sum(first, second int) int {
	sum := first + second
	return sum
}

func SumAndMultiply(first, second int) (int, int) {
	return first + second, first * second
}

func namedSumAndMultiply(first, second int) (sum int64, multiply int64) {
	sum = int64(first + second)
	multiply = int64(first) * int64(second)
	return // или return sum, multiply
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

func createDivider(divider int /* 2 */) func(x int) int {
	return func(x int) int {
		return x / divider /* 2 */
	}
}

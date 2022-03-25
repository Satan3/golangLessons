package main

import (
	"fmt"
)

func main() {
	// defer block
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(sum(2, 3))
	//deferValues()

	// goroutines block
	//runtime.GOMAXPROCS(1)
	//fmt.Println(runtime.NumCPU())

	//go showNumbers(100)

	//runtime.Gosched()
	//time.Sleep(time.Second)
	fmt.Println("exit")

	makePanic()
}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y
	return
}

func deferValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("first", i)
	}
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("second", i)
		}()
	}

	for i := 0; i < 10; i++ {
		k := i
		defer func() {
			fmt.Println("third", k)
		}()
	}
	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println("fourth", k)
		}(i)
	}
}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	panic("some panic")
	fmt.Println("Unreachable code")
}

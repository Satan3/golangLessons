package main

import (
	"fmt"
	"time"
)

func main() {
	//nilChannel()
	//unbufferedChannel()
	//bufferedChannel()
	//forRange()
}

func nilChannel() {
	var nilChannel chan int
	fmt.Printf("Len: %d Cap: %d\n", len(nilChannel), cap(nilChannel))

	// write to nil channel blocks forever
	//nilChannel <- 1

	// read from nil channel blocks forever
	//<-nilChannel

	// closing nil channel will raise a panic
	//close(nilChannel)
}

func unbufferedChannel() {
	unbufferedChannel := make(chan int)
	fmt.Printf("Len: %d Cap: %d\n", len(unbufferedChannel), cap(unbufferedChannel))

	// blocks until smb reads
	//unbufferedChannel <- 1

	// blocks until smb write
	//<-unbufferedChannel

	// blocks on reading then write
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		// <-chanForWriting
		unbufferedChannel <- 1
	}(unbufferedChannel)

	value := <-unbufferedChannel
	fmt.Println(value)

	// blocks on writing then read
	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-chanForReading
		fmt.Println(value)
	}(unbufferedChannel)

	unbufferedChannel <- 2

	// close channel
	/*go func() {
		time.Sleep(time.Millisecond * 500)
		close(unbufferedChannel)
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println(<-unbufferedChannel)
	}()
	// write to closed chan panics
	unbufferedChannel <- 3*/

	// close of closed panics
	//close(unbufferedChannel)
	//close(unbufferedChannel)
}

func bufferedChannel() {
	bufferedChannel := make(chan int, 2)
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	// doesn`t block while buffer not full
	bufferedChannel <- 1
	bufferedChannel <- 2

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	// blocks to write, buffer is full
	//bufferedChannel <- 3

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	// blocks to read, buffer is empty, no writers
	//fmt.Println(<-bufferedChannel)
}

func forRange() {
	bufferedChannel := make(chan int, 3)

	numbers := []int{5, 6, 7, 8}

	// show all elements with for

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for {
		v, ok := <-bufferedChannel
		fmt.Println(v, ok)

		if !ok {
			break
		}
	}

	// show with for range buffered
	bufferedChannel = make(chan int, 3)

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for v := range bufferedChannel {
		fmt.Println("buffered", v)
	}

	unbufferedChannel := make(chan int)
	go func() {
		for _, num := range numbers {
			unbufferedChannel <- num
		}
		close(unbufferedChannel)
	}()

	for value := range unbufferedChannel {
		fmt.Println("unbuffered", value)
	}
}

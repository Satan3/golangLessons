package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//baseSelect()
	//gracefulShutdown()
}

func baseSelect() {
	bufferedChan := make(chan string, 3)
	bufferedChan <- "first"

	select {
	case str := <-bufferedChan:
		fmt.Println("read", str)
	case bufferedChan <- "second":
		fmt.Println("write", <-bufferedChan, <-bufferedChan)
	}

	unbufChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		unbufChan <- 1
	}()

	select {
	case bufferedChan <- "third":
		fmt.Println("unblocking writing")
	case val := <-unbufChan:
		fmt.Println("blocking reading", val)
	case <-time.After(time.Millisecond * 1500):
		fmt.Println("time`s up")
	default:
		fmt.Println("default case")
	}

	resultChan := make(chan int)
	timer := time.After(time.Second) // timer outside loop

	go func() {
		defer close(resultChan)

		for i := 1; i <= 1000; i++ {

			select {
			case <-timer:
				fmt.Println("time`s up")
				return
			default:
				//time.Sleep(time.Nanosecond)
				resultChan <- i
			}
		}
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	select {
	case <-timer:
		fmt.Println("time`s up")
		return
	case sig := <-sigChan:
		fmt.Println("Stopped by signal:", sig)
		return
	}
}

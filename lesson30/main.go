package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//output()
	//input()
	args()
}

func output() {
	count, err := fmt.Println("some text")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

	count, err = fmt.Fprintln(os.Stdout, "stdout")
	if err != nil {
		log.Fatal(err)
	}

	count, err = fmt.Fprintln(os.Stderr, "stderr")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err = fmt.Fprintln(file, "hello\nworld")
	if err != nil {
		log.Fatal(err)
	}
}

func input() {
	var (
		text  string
		text2 string
	)

	count, err := fmt.Scan(&text, &text2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(text, text2, count)

	count, err = fmt.Fscan(os.Stdin, &text, &text2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(text, text2, count)

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err = fmt.Fscanln(file, &text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(text, count)
}

func args() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

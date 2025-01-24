package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	//createFile()
	//writeTo()
	//readFrom()
	//ioCopy()
	bufioWrite()
	//createFileBufio()
}

func createFile() {
	start := time.Now()

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	var i int

	for i < 100 {
		if _, err = file.WriteString(fmt.Sprintf("%d\n", i)); err != nil { // check syscalls
			log.Fatal(err)
		}

		i++
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}

func writeTo() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err := file.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func readFrom() {
	file, err := os.Create("test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	file2, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file2.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err := file.ReadFrom(file2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func ioCopy() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//count, err := io.Copy(os.Stdout, file)
	count, err := io.CopyBuffer(os.Stdout, file, make([]byte, 1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println(count)
}

func bufioWrite() {
	var err error

	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		if err == nil {
			if err := writer.Flush(); err != nil {
				log.Fatal(err)
			}
		}
	}()

	// do some logic 1

	// write log 1
	if _, err := writer.WriteString("first action done\n"); err != nil {
		log.Fatal(err)
	}

	// do some logic 2

	// write log 2
	if _, err := writer.WriteString("second action done\n"); err != nil {
		log.Fatal(err)
	}

	err = errors.New("some error")
}

func createFileBufio() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	file, err := os.Create("test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	writer := bufio.NewWriter(file)
	defer func() {
		if err = writer.Flush(); err != nil {
			log.Fatal(err)
		}
	}()

	var i int

	for i < 100 {
		if _, err = writer.WriteString(fmt.Sprintf("%d\n", i)); err != nil { // check syscalls
			log.Fatal(err)
		}

		i++
	}
}

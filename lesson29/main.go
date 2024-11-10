package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//io.Reader
	//io.Writer

	//simpleReader()
	//rowsReader()

	//simpleWriter()

	osFile()
}

func simpleReader() {
	buf := make([]byte, 1)

	reader := NumsReader{nums: "1,2,3,4,5,6.7=,g8"}

	count, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println(string(buf), count)
}

type NumsReader struct {
	nums string
}

func (r NumsReader) Read(p []byte) (n int, err error) {
	var count int
	for i := 0; i < len(r.nums); i++ {
		if r.nums[i] >= '0' && r.nums[i] <= '9' {
			p[count] = r.nums[i]
			count++
		}
	}

	return count, io.EOF
}

func rowsReader() {
	rowsReader := RowsReader{text: "первая строка\nвторая строка\nтретья строка"}
	var (
		err   error
		count int
	)

	row := make([]byte, 100)
	for err != io.EOF {
		count, err = rowsReader.Read(row)
		fmt.Println(string(row), count)
	}
}

type RowsReader struct {
	text string
}

func (r *RowsReader) Read(p []byte) (n int, err error) {
	var i int
	for i = 0; i < len(r.text); i++ {
		if r.text[i] == '\n' {
			r.text = r.text[i+1:]
			break
		}

		p[i] = r.text[i]

		if i == len(r.text)-1 {
			r.text = ""
			return i + 1, io.EOF
		}
	}

	return i + 1, nil
}

func simpleWriter() {
	nums := []byte{'1', ',', '2', '=', '3', '.'}

	writer := NumsWriter{storedNums: make([]byte, 10)}
	count, err := writer.Write(nums)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(writer.storedNums), count)
}

type NumsWriter struct {
	storedNums []byte
}

func (w NumsWriter) Write(p []byte) (n int, err error) {
	var count int

	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			w.storedNums[count] = p[i]
			count++
		}
	}

	return count, nil
}

func osFile() {
	newFile, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	n, err := newFile.WriteString("hello world")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	err = newFile.Close()
	if err != nil {
		log.Fatal()
	}

	file, err := os.Open("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 100)

	n, err = file.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf), n)

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	file, err = os.OpenFile("new.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if n, err = file.WriteString("\nhello world 2"); err != nil {
		log.Fatal(err)
	}

	// check file
}

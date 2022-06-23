package main

import (
	"fmt"
	fiber2 "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println(fiber2.New())

	fmt.Println("Without dependencies")
}

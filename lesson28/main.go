package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	fiber2 "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println(fiber.New())
	fmt.Println(fiber2.New())

	fmt.Println("Without dependencies")
}

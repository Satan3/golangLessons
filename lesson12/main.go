package main

import (
	"fmt"
	"time"
)

type OurType string

// 1. Можем так сделать, ибо тип в этом же пакете
func (t OurType) Hello() {
	fmt.Println("Hello")
}

// Так не можем, ибо Time находится в другом пакете
/*
func (t time.Time) HelloTime() {
	fmt.Println("Hello time")
}*/

// 2. Так не можем, ибо тип не должен быть указателем
type pInt *int

/*func (pInt) Test() {
	fmt.Println("Test")
}
*/

// 3. Интерфейсный тип
type Tester interface {
	Hello()
}

/*func (t Tester) Test() {
	fmt.Println("Test")
}*/

// 4. Builtin тип
/*func (i int64)Test() {
	fmt.Println("Test")
}*/

func main() {
	//definition()
	rules()
}

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("Периметр фигуры: %d \n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func (s Square) WrongScale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func definition() {
	square := Square{Side: 4}
	pSquare := &square

	square2 := Square{Side: 2}

	square.Perimeter() // Square{Side:4}
	square2.Perimeter()

	pSquare.Scale(2) // &Square{Side:4}

	pSquare.Perimeter() // (*pSquare).Perimeter() Square{Slide:8}
	square.Scale(2)     // (&square).Scale
	pSquare.Perimeter() // (*pSquare).Perimeter()

	square.WrongScale(2)
	square.Perimeter()
}

func rules() {
	// 1. Тип должен быть объявлен в текущем пакете
	// 5. Является объявленным типом

	now := time.Now()
	fmt.Printf("%T, %#v \n", now, now)

	var ourString OurType = "hello"
	fmt.Printf("%T, %#v \n", ourString, ourString)

	// 2. Тип не должен быть указателем
	var pInt pInt
	fmt.Printf("%T, %#v \n", pInt, pInt)

	// 3. Интерфейсный тип
	var stringTwo Tester
	fmt.Printf("%T, %#v \n", stringTwo, stringTwo)

	// 4. builtin тип
	var builtinType int64 = 1
	fmt.Printf("%T, %#v \n", builtinType, builtinType)
}

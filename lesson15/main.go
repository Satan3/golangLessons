package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	arrays()
	slices()
}

func arrays() {
	var intArr [3]int
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	intArr[0] = 5
	intArr[1] = 6
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	people := [2]Person{
		{
			Age:  30,
			Name: "Katy",
		},
		{
			Age:  23,
			Name: "John",
		},
	}
	fmt.Printf("Type: %T Value: %#v\n", people, people)

	stringsArr := [...]string{"First", "Second", "Third", "Fourth"}
	fmt.Printf("Type: %T Value: %#v\n", stringsArr, stringsArr)

	fmt.Printf("Length: %d Capacity: %d\n", len(stringsArr), cap(stringsArr))

	for i := 0; i < len(stringsArr); i++ {
		fmt.Printf("Index: %d Value: %s\n", i, stringsArr[i])
	}

	for inx, value := range stringsArr {
		fmt.Printf("Index: %d Value: %s\n", inx, value)
	}

	for ind := range intArr {
		fmt.Printf("Value: %d\n", ind)
	}

	newIntArr := changeArray(intArr)
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)
	fmt.Printf("Type: %T Value: %#v\n", newIntArr, newIntArr)

	// out of bounds error
	//fmt.Println(newIntArr[4])
}

func changeArray(arr [3]int) [3]int {
	arr[2] = 3
	return arr
}

func slices() {
	var defaultSlice []int
	fmt.Printf("Type: %T Value: %#v\n", defaultSlice, defaultSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(defaultSlice), cap(defaultSlice))
	// out of range
	//defaultSlice[0] = 1

	stringSliceLiteral := []string{"First", "Second"}
	fmt.Printf("Type: %T Value: %#v\n", stringSliceLiteral, stringSliceLiteral)
	fmt.Printf("Length: %d Capacity: %d\n", len(stringSliceLiteral), cap(stringSliceLiteral))

	sliceByMake := make([]int, 0, 5)
	fmt.Printf("Type: %T Value: %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5)
	fmt.Printf("Type: %T Value: %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 6)
	fmt.Printf("Type: %T Value: %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	for ind, value := range sliceByMake {
		fmt.Printf("Index: %d Value: %d\n", ind, value)
	}
}

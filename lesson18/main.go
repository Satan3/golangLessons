package main

import "fmt"

type User struct {
	Id   int64
	Name string
}

func main() {
	// default value
	var defaultMap map[int64]string

	fmt.Printf("Type: %T Value: %#v\n", defaultMap, defaultMap)
	fmt.Printf("Len: %d\n\n", len(defaultMap))

	// slice by make
	mapByMake := make(map[string]string)
	fmt.Printf("Type: %T Value: %#v\n\n", mapByMake, mapByMake)

	// slice by make with cap
	mapByMakeWithCap := make(map[string]string, 3)
	fmt.Printf("Type: %T Value: %#v\n\n", mapByMakeWithCap, mapByMakeWithCap)

	// slice by literal
	mapByLiteral := map[string]int{"Vasya": 18, "Dima": 20}
	fmt.Printf("Type: %T Value: %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Len: %d\n\n", len(mapByLiteral))

	// slice by new
	mapWithNew := *new(map[string]string)
	fmt.Printf("Type: %T Value: %#v\n\n", mapWithNew, mapWithNew)

	// insert value
	mapByMake["First"] = "Vasya"
	fmt.Printf("Type: %T Value: %#v\n", mapByMake, mapByMake)
	fmt.Printf("Len: %d\n\n", len(mapByMake))

	// update value
	mapByMake["First"] = "Petya"
	fmt.Printf("Type: %T Value: %#v\n", mapByMake, mapByMake)
	fmt.Printf("Len: %d\n\n", len(mapByMake))

	// get map value
	fmt.Println(mapByLiteral["Vasya"])

	// get map default value
	fmt.Println(mapByLiteral["Second"])

	// check value existence
	value, ok := mapByLiteral["Second"]
	fmt.Printf("Value: %d IsExist: %t\n\n", value, ok)

	// delete value
	delete(mapByMake, "First")
	fmt.Printf("Type: %T Value: %#v\n\n", mapByMake, mapByMake)

	// map iteration
	mapForIteration := map[string]int{"First": 1, "Second": 2, "Third": 3, "Fourth": 4}

	for key, val := range mapForIteration {
		fmt.Printf("Key: %s Value: %d\n", key, val)
	}

	// unique values
	users := []User{
		{
			Id:   1,
			Name: "Vasya",
		},
		{
			Id:   45,
			Name: "Petya",
		},
		{
			Id:   57,
			Name: "John",
		},
		{
			Id:   45,
			Name: "Petya",
		},
	}

	uniqueUsers := make(map[int64]struct{}, len(users))

	for _, user := range users {
		if _, ok := uniqueUsers[user.Id]; !ok {
			uniqueUsers[user.Id] = struct{}{}
		}
	}
	fmt.Printf("Type: %T Value: %#v\n", uniqueUsers, uniqueUsers)

	// find by value
	usersMap := make(map[int64]User, len(users))
	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}

	fmt.Println(findInSlice(57, users))
	fmt.Println(findInMap(57, usersMap))
}

func findInSlice(id int64, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}

	return nil
}

func findInMap(id int64, usersMap map[int64]User) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}

	return nil
}

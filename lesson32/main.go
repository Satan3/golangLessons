package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	//withoutTags()
	//withTags()
	//unknownJson()
	//custom()
	//encoderDecoder()
}

func withoutTags() {
	person := Person{
		Name:    "Vasya",
		Age:     25,
		Married: false,
		Hobbies: []string{
			"programming",
			"sport",
		},
		Address: Address{
			Country: "RU",
			City:    "Moscow",
		},
	}

	bytes, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))

	var resultPerson Person

	if err = json.Unmarshal(bytes, &resultPerson); err != nil {
		log.Fatal(err)
	}

	fmt.Println(person)
}

func withTags() {
	person := PersonWithTags{
		Name:    "Vasya",
		Age:     25,
		Married: false,
		Hobbies: nil,
		Address: Address{
			Country: "RU",
			City:    "Moscow",
		},
	}

	bytes, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))

	var resultPersonWithTags PersonWithTags

	if err = json.Unmarshal(bytes, &resultPersonWithTags); err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultPersonWithTags)
}

func unknownJson() {
	request := `{"first":1, "second":[1,2]}`

	result := make(map[string]any)

	if err := json.Unmarshal([]byte(request), &result); err != nil {
		log.Fatal(err)
	}

	for k, v := range result {
		fmt.Printf("Ключ: %s, Значение: %v \n", k, v)
	}
}

func custom() {
	person := &CustomPerson{
		Name:    "Vasya",
		Age:     25,
		Married: false,
		Hobbies: []string{},
		Address: Address{
			Country: "RU",
			City:    "Moscow",
		},
	}

	bytes, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))

	var resultPerson CustomPerson

	if err = json.Unmarshal(bytes, &resultPerson); err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultPerson)
}

func encoderDecoder() {
	file, err := os.Open("./lesson32/example.json")
	if err != nil {
		log.Fatal(err)
	}

	var person PersonWithTags

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&person); err != nil {
		log.Fatal(err)
	}

	fmt.Println(person)

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	// encoding
	person.Name = "Petya"

	file, err = os.OpenFile("./lesson32/example.json", os.O_WRONLY|os.O_TRUNC, 777)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = json.NewEncoder(file).Encode(person); err != nil {
		log.Fatal(err)
	}
}

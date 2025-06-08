package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int8
	Married bool
	Hobbies []string
	Address Address
}

type Address struct {
	Country string
	City    string
}

type PersonWithTags struct {
	Name    string   `json:"name"`
	Age     int8     `json:"age"`
	Married bool     `json:"married"`
	Hobbies []string `json:"hobbies,omitempty"`
	Address Address  `json:"-"`
}

type CustomPerson struct {
	Name    string
	Age     int8
	Married bool
	Hobbies []string
	Address Address
}

func (p *CustomPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string   `json:"name"`
		Age     int8     `json:"age"`
		Married bool     `json:"married"`
		Hobbies []string `json:"hobbies,omitempty"`
		Address struct {
			Country string `json:"country"`
		} `json:"address"`
	}{
		Name:    p.Name,
		Age:     p.Age,
		Married: p.Married,
		Hobbies: p.Hobbies,
		Address: struct {
			Country string `json:"country"`
		}{
			Country: p.Address.Country,
		},
	})
}

func (p *CustomPerson) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name    string   `json:"name"`
		Age     int8     `json:"age"`
		Married bool     `json:"married"`
		Hobbies []string `json:"hobbies,omitempty"`
		Address struct {
			Country string `json:"country"`
		}
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	fmt.Println(temp)

	p.Name = temp.Name
	p.Age = temp.Age
	p.Married = temp.Married
	p.Hobbies = temp.Hobbies
	p.Address = Address{
		Country: temp.Address.Country,
		City:    "",
	}

	return nil
}

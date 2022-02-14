package main

import "fmt"

type Builder interface {
	Build()
}

type Person struct {
	Name string
	Age  int
}

type WorkExperience struct {
	Name string
	Age  int
}

func (p Person) printName() {
	fmt.Println(p.Name)
}

type WoodBuilder struct {
	Person
	//Name string
	//WorkExperience
}

/*func (wb WoodBuilder) printName() {
	fmt.Println(wb.Name)
}*/

func (wb WoodBuilder) Build() {
	fmt.Println("Строю дом из дерева")
}

type BrickBuilder struct {
	Person
}

func (bb BrickBuilder) Build() {
	fmt.Println("Строю из кирпича")
}

type Building struct {
	Builder
	Name string
}

func main() {
	//explanation()
	usecase()
}

/*func explanation() {
	//builder := WoodBuilder{Person{Name: "Вася", Age: 30}}
	//builder := WoodBuilder{Person{Name: "Вася", Age: 30}, "Боб"}
	builder := WoodBuilder{
		Person{Name: "Вася", Age: 30},
		"Боб",
		WorkExperience{Name: "Таксист", Age: 3}}
	fmt.Printf("Type: %T Value: %#v\n", builder, builder)

	// shorthands
	fmt.Println(builder.Person.Age)
	fmt.Println(builder.WorkExperience.Age)

	//shadowing
	fmt.Println(builder.Name)
	fmt.Println(builder.Person.Name)

	builder.printName()
}*/

func usecase() {
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			Name: "Вася",
			Age:  40,
		}},
		Name: "Деревянная изба",
	}
	woodenBuilding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{
			Person{
				Name: "Петя",
				Age:  30,
			},
		},
		Name: "Кирпичный дом",
	}
	brickBuilding.Build()
}

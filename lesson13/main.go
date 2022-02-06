package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name)
}

func (h Human) writeCode() {
	fmt.Println("Человек пишет код")
}

type Duck struct {
	Name, Surname string
}

func (d Duck) Run() string {
	return "Утка бегает"
}

func (d Duck) Swim() string {
	return "Утка плавает"
}

func (d Duck) Fly() string {
	return "Утка летает"
}

func main() {
	//interfaceValues()
	typeAssertionAndPolymorphism()
}

func interfaceValues() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	if runner == nil {
		fmt.Println("Runner is nil")
	}
	//runner = int64(1)
	//runner.Run()

	var unnamedRunner *Human
	fmt.Printf("Type: %T Value: %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("Runner is nil")
	}

	namedRunner := &Human{Name: "Джек"}
	fmt.Printf("Type: %T Value: %#v\n", namedRunner, namedRunner)

	runner = namedRunner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	// empty interface{}

	var emptyInterface interface{} = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = runner
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = int64(1)
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = true
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)
}

func typeAssertionAndPolymorphism() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	john := &Human{"John"}
	runner = john
	polymorphism(john)
	typeAssertion(john)

	donald := &Duck{Name: "Donald", Surname: "Duck"}
	runner = donald
	polymorphism(donald)
	typeAssertion(donald)
}
func polymorphism(runner Runner) {
	fmt.Println(runner.Run())
}

func typeAssertion(runner Runner) {
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if human, ok := runner.(*Human); ok {
		fmt.Printf("Type: %T Value: %#v\n", human, human)
		human.writeCode()
	}
	if duck, ok := runner.(*Duck); ok {
		fmt.Printf("Type: %T Value: %#v\n", duck, duck)
		fmt.Println(duck.Fly())
	}

	switch v := runner.(type) {
	case *Human:
		fmt.Println(v.Run())
	case *Duck:
		fmt.Println(v.Swim())
	default:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
	}
}

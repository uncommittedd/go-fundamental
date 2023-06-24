package main

import "fmt"

type HasName interface {
	GetName() string
}

type People struct {
	Name string
}

type Animal struct {
	Name string
}

func (person People) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayWord(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

func main() {
	mario := People{Name: "mario"}
	sayWord(mario)

	var cat Animal
	cat.Name = "kitten"

	sayWord(cat)
}

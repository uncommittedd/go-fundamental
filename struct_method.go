package main

import "fmt"

type Hero struct {
	name, address string
	age           int
}

func (hero Hero) attack(name string) {
	fmt.Println(hero.name, "attack", name)
}

func (hero Hero) heal(name string) {
	fmt.Printf("%s heal %s", hero.name, name)
}

func main() {
	io := Hero{
		name:    "Io",
		address: "MMo",
		age:     2000,
	}
	io.attack("Glo")
	io.heal("Clint")
}

package main

import "fmt"

func sayHello(name string) string {
	result := fmt.Sprintf("hello %s", name)
	return result
}

func lol(toSomeone string) string {
	return "laughing out loud, " + toSomeone
}

func main() {
	fmt.Println(sayHello("mario"))
	fmt.Println(lol("john"))
}

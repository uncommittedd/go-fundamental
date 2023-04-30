package main

import "fmt"

func saySomething(name string) string {
	return "hay " + name
}

func main() {
	goodMorning := saySomething
	fmt.Println(goodMorning("mario"))
}

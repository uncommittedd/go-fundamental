package main

import "fmt"

func endRuntime() {
	fmt.Println("The runtime is stopped")
	errorMessage := recover()
	fmt.Println("Error :", errorMessage)
}

func runRuntime(isError bool) {
	defer endRuntime()
	if isError {
		panic("Error, because ur noob")
	}
	fmt.Println("This runtime is on EOD")
}

func main() {
	runRuntime(true)
	fmt.Println("recover complete")
}

/**
panic, if this function called the runtime will be stopped, this schema like throw Error
*/
package main

import "fmt"

func endApp() {
	fmt.Println("Application stopped")
}

func runApp(isError bool) {
	defer endApp()
	if isError {
		panic("Application Error")
	}
	fmt.Println("Application Running")
}

func main() {
	runApp(true)
}

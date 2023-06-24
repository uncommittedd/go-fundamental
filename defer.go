/**
This defer function will called on end of run of a function.
if error on code this function will be called
the
*/
package main

import "fmt"

func logging() {
	fmt.Println("This log is called")
}

func runApplication(number int16) {
	defer logging()
	result := 10 / number
	fmt.Println("App Started")
	fmt.Println(result)
}

func main() {
	runApplication(0)
}

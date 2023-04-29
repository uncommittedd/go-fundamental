package main

import "fmt"

func main() {
	counter := 0
	for counter <= 6 {
		fmt.Println(counter)
		counter++
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("iteration %d \n", i)
	}

	fullName := []string{"Mario", "Haryzal"}
	for i := 0; i < len(fullName); i++ {
		fmt.Println(fullName[i])
	}

	for i, v := range fullName {
		fmt.Println("index", i, "=", v)
	}

	fullNameMap := make(map[string]string)
	fullNameMap["firstName"] = "Mario"
	fullNameMap["lastName"] = "Haryzal"
	for k, v := range fullNameMap {
		fmt.Println(k, "=", v)
	}
}

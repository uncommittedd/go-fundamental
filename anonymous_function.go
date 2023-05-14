package main

import "fmt"

type BlackList func(string) bool

func checkUsername(username string, isBlocked BlackList) {
	if isBlocked(username) {
		fmt.Println("cannot login ur account is invalid")
	} else {
		fmt.Println("welcome " + username)
	}
}

func main() {
	checkUsername("Mario", func(name string) bool {
		return name == "Shit"
	})
}

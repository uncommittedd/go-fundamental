package main

import "fmt"

func main() {
	roles := []string{"mid", "exp", "gold", "roam", "jungle"}

	if length := len(roles); length == 3 {
		fmt.Println("this game is MOBA")
	} else {
		fmt.Println("idk, maybe action game")
	}
}

package main

import "fmt"

func main() {
	var hero string = "roam"

	if hero == "roam" {
		fmt.Printf("this hero is %v type", hero)
	} else if hero == "gold" {
		fmt.Printf("this hero is %v type", hero)
	} else if hero == "exp" {
		fmt.Printf("this hero is %v type", hero)
	} else if hero == "jungle" {
		fmt.Printf("this hero is %v type", hero)
	} else {
		fmt.Println("maybe this hero is minion")
	}
}

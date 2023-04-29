package main

import "fmt"

func main() {
	items := map[string]string{
		"early":  "boots",
		"middle": "blade armor",
	}

	fmt.Println(items)
	fmt.Println(items["early"])
	fmt.Println(items["middle"])
	fmt.Println(len(items))
	items["late"] = "immortal"
	fmt.Println(items)
	delete(items, "late")
	fmt.Println(items)

	var heros map[string]string = make(map[string]string)
	heros["mid"] = "mage"
	heros["exp"] = "fighter"
	fmt.Println(heros)
}

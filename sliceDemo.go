package main

import "fmt"

func main() {
	months := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	sliceOne := months[4:7]
	sliceTwo := months[6:9]

	fmt.Println(months)
	fmt.Println(sliceOne)
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))
	fmt.Println(sliceTwo)
	fmt.Println(len(sliceTwo))
	fmt.Println(cap(sliceTwo))
	fmt.Println("====================================")

	/*
		===========================================
	*/
	//changes on array will make slice changes too
	// months[5] = "change"
	// fmt.Println(sliceOne)

	// same if change slice will be changed too on array referenced
	sliceOne[0] = "change too"
	fmt.Println(sliceOne)
	fmt.Println(months)
	fmt.Println("==============================")

	days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	newFormatDay := append(days[2:4], "new day")
	fmt.Println(cap(newFormatDay))
	fmt.Println(newFormatDay)
	fmt.Println(len(newFormatDay))
	fmt.Println(cap(newFormatDay))

	sliceThree := append(newFormatDay, "Mario")
	sliceThree[0] = "newest"
	fmt.Println(sliceThree)
	fmt.Println(newFormatDay)
	fmt.Println(days)
}

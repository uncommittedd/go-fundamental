package main

import "fmt"

func getHero() (firstRecommend, secondRecommend, thirdRecommend string) {
	return "Mathilda", "Glo", "Angela"
}

func main() {
	firstHero, secondHero, _ := getHero()
	fmt.Println(firstHero, secondHero)
}

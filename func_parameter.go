package main

import "fmt"

type Filter func(string) string

func saySomeQuotes(quote string, filterQuotes Filter) {
	quoteFiltered := filterQuotes(quote)
	fmt.Println(quoteFiltered)
}

func filter(word string) string {
	if word == "Dog" || word == "shit" {
		return "..."
	} else {
		return word
	}
}

func main() {
	saySomeQuotes("Mario", filter)
}

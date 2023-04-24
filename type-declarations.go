package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var (
		noKtpMario    NoKtp = "2342423423"
		marriedStatus Married
	)
	noKtpMario = "342342525"
	marriedStatus = false

	fmt.Println(noKtpMario)
	fmt.Println(marriedStatus)

}

package main

import "fmt"

func main() {

	type NoKTP string
	type married bool

	var noKtpChristian NoKTP = "1234567890"
	var marriedStatus married = true

	fmt.Println(noKtpChristian)
	fmt.Println(marriedStatus)
}

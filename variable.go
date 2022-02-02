package main

import "fmt"

func main() {
	var name string

	name = "Christian Eka"
	fmt.Println(name)
	name = "Kenneth Kennard"
	fmt.Println(name)

	var friendName = "kenneth"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	country := "indonesia" // Deklarasi awal
	fmt.Println(country)

	//country := "America" // ini Error karna := hanya untuk deklarasi awal
	country = "America"

	var (
		firstName = "Christian"
		lastName  = "Eka"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}

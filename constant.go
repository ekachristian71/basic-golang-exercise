package main

import "fmt"

func main() {

	const firstName string = "Christian"
	const lastName = "Eka"
	const value = 1000

	fmt.Println(lastName)
	fmt.Println(value)

	//firstName = "haha" Error karna Constant tidak bisa diubah valuenya

	const (
		contoh1 string = "kenneth"
		contoh2        = "kennard"
		nilai          = 1000
	)
	fmt.Println(contoh1)
	fmt.Println(contoh2)
	fmt.Println(nilai)
}

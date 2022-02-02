package main

import "fmt"

func main() {

	var name1 = "christian"
	var name2 = "christian"
	var result bool = name1 == name2
	fmt.Println("Result = ", result) // result nya TRUE

	var name3 = "dhristian"
	var name4 = "christian"
	var result2 bool = name3 > name4
	fmt.Println("Result 2 = ", result2) // resultnya TRUE karena karakter d pada dchristian lebih dari c pada christian

	var value1 = 100
	var value2 = 200

	fmt.Println("Result 3")
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}

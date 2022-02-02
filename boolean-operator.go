package main

import "fmt"

func main() {
	var valueEnd = 80
	var absence = 80

	var gradValue = valueEnd >= 80
	var gradAbsence = absence >= 80

	var finalGrad = gradValue && gradAbsence

	fmt.Println(finalGrad)
}

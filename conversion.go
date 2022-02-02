package main

import "fmt"

func main() {

	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	//var nilai8 int8 = int8(nilai32) kena limit maksimum

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var name = "Christian"
	var c byte = name[0]
	var cString = string(c)

	fmt.Println(name)
	fmt.Println(cString)
}

package main

import "fmt"


var someVariable = "Hello"

func main() {
	
	// String Declaration 
	fmt.Println("--------------------------------")
	fmt.Println("String Declaration")
	var nameOne string = "Irshad"
	var nameTwo = "Ibrahim"
	nameThree := "Kadungamkunnath"
	var startOne string
	startOne = "Name is : "
	fmt.Println(startOne, nameOne, nameTwo, nameThree)

	fmt.Println(someVariable)

	// Integer  signed
	fmt.Println("--------------------------------")
	fmt.Println("Integer Declaration")
	var n1 int8 = 127 //max 127
	var n2 int16 =  -123
	var n3 int32 = 46546464
	var n4 int64 = -4646545646

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

	// unsigned integer
	var n5 uint8 = 127 //max 127
	var n6 uint16 =  123
	var n7 uint32 = 46546464
	var n8 uint64 = 4646545646

	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)

	
	fmt.Println("--------------------------------")
	fmt.Println("Float Declaration")

	var n9 float32 = 25.65
	var n10 float64 = 888.454

	fmt.Println(n9)
	fmt.Println(n10)

	 
}
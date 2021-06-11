package main

import "fmt"

func main() {
	fmt.Print("Hello \n")	

	age := 10
	name := "Irshad"

	fmt.Println("name is",name, "age is", age)

	fmt.Printf("name is %v and age is %v \n", name, age) // formated 

	fmt.Printf("score is : %f \n", 10.44)

	fmt.Printf("type of age is %T \n", age)

	// #######################################################################
	//save formatted string
	var str = fmt.Sprintf("Name is %v and age is %v", name, age)
	println(str)


}
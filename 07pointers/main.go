package main

import "fmt"

func main() {
	fmt.Println("pointer")
	// Just create pointer
	// var ptr *int
	// fmt.Println("Pointer val : ", ptr)

	//Creating pointer with reference some value
	mynumber := 23
	//this will store memory address || If print with *ptr this will display value allocated to memory address
	var ptr = &mynumber
	fmt.Println("Variable val : ", mynumber) //here we get variable value
	fmt.Println("Pointer val : ", ptr)       //here we get memory address
	fmt.Println("Pointer val : ", *ptr)      //here we get value wich store at this memory address

	mynumber = mynumber + 2
	*ptr = mynumber + 2
	fmt.Println("Variable val : ", mynumber) //here we get variable value
	fmt.Println("Pointer val : ", ptr)       //here we get memory address
	fmt.Println("Pointer val : ", *ptr)      //here we get value wich store at this memory address
}

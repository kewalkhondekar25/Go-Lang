package main

import "fmt"

func main() {

	fmt.Println("welcome to pointers")

	var ptr *int
	fmt.Println("value of pointer :", ptr)

	//var
	myNumber := 23

	//reference to pointer
	myPointer := &myNumber

	//pointer is reference to direct memory location
	fmt.Println("value of actual pointer is :", myPointer)

	//* means what value it holds
	fmt.Println("value of actual pointer is :", *myPointer)
}

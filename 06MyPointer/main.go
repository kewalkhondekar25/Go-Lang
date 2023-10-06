package main

import "fmt"

func main() {

	fmt.Println("welcome to pointers")

	var ptr *int
	fmt.Println("value of pointer :", ptr)

	myNumber := 23
	myPointer := &myNumber
	fmt.Println("value of actual pointer is :", myPointer)
	fmt.Println("value of actual pointer is :", *myPointer)
}

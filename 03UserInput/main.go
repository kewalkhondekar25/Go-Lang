package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	//NewReader return a new Reader && os reads from I/O
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our consultation service: ")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us", input)
}

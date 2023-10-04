package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to Conversion")
	fmt.Println("Please enter a number between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("You have Enter: ", input)
	fmt.Printf("type of input is %T \n", input)

	//parsing && triming spaces
	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your entered value: ", rating+1)
	}
}

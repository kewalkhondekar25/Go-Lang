package main

import "fmt"

func main() {

	fmt.Println("Welcome to introduction of array")

	var fruitLists [4]string

	fruitLists[0] = "apple"
	fruitLists[1] = "grapes"
	fruitLists[3] = "peach"

	fmt.Println("fruit list is :", fruitLists)
	//output
	//fruit list is : [apple grapes  peach]

	var veggiesList = [3]string{"potato", "broccoli", "pumpkin"}

	fmt.Println("veggies list is :", veggiesList)
	//output
	//veggies list is : [potato broccoli pumpkin]

}

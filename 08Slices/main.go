package main

import "fmt"

func main() {

	fmt.Println("weelcome to slices")

	//in slices, don't mention number of values unlike arrays
	var codeLangList = []string{"go", "javascrit", "python"}

	fmt.Println("values of slices are :", codeLangList)
	fmt.Printf("type of code LangList is %T \n :", codeLangList)
}

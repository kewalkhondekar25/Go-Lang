package main

import "fmt"

func main() {

	fmt.Println("weelcome to slices")

	//in slices, don't mention number of values unlike arrays
	var codeLangList = []string{"go", "javascrit", "python"}

	fmt.Println("values of slices are :", codeLangList)
	fmt.Printf("type of code LangList is %T \n :", codeLangList)

	//append data in slices
	codeLangList = append(codeLangList, "c", "java")
	fmt.Println("final codeLangList :", codeLangList)

	//slice the slices
	//[n:n] starting index to land index which is range non-inclusive
	codeLangList = append(codeLangList[1:3])
	fmt.Println("final codeLandList is :", codeLangList)

	codeLangList = append(codeLangList[:3])
	fmt.Println("final codeLangList values :", codeLangList)
}

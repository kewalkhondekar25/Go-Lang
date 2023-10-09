package main

import (
	"fmt"
	"sort"
)

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

	//make need to be define type of data & no of values
	highScores := make([]int, 4)

	highScores[0] = 25
	highScores[1] = 44
	highScores[2] = 34
	highScores[3] = 99

	//make will gave default allocation of memory; append will reallocate memory
	highScores = append(highScores, 787, 350, 777)

	fmt.Println("highscores values are :", highScores)

	//sort pkg sorts string, int etc...
	sort.Ints(highScores)
	fmt.Println("sorted values are :", highScores)
}

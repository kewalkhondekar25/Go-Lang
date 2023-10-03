//variables, types, const

package main

import "fmt"

//walrus operator not allow outside methode
//jwtToken := 123987456

//allowed
var jwtTOKEN = 123456789

//public variable (caoitalize)
const LoginToken string = "gibbrisshh"

func main() {
	var userName string = "kewal"
	fmt.Println(userName)
	fmt.Printf("Variable is of type : %T \n", userName)

	// var userName2 string = ""
	// fmt.Println(userName2)
	// fmt.Printf("Variable is of type : %T \n", userName2)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	//unit8 (0 to 255)
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T \n", smallValue)

	var smallFloat float32 = 25.12345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var integer int
	fmt.Println(integer)
	fmt.Printf("variable is of type : %T \n", integer)

	//implicit type
	var website = "kewalkhondekar.dev"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	//no var type {walrus operator}
	numberOfPeople := 2500
	fmt.Println(numberOfPeople)
	fmt.Printf("Variable is of type : %T \n", numberOfPeople)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)
}

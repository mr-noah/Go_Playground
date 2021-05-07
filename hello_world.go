//Noah Lambaria (Notes on Go -> for internship preparation)

package main

import "fmt"

//can also declare variables at the package level; however,
//this cant use the ":=" syntax
var test int = 32

//if variable is lowercase, variable is scoped to the package,
//else if uppercase then variable is globally visible

//clean code technique
var (
	stringTest string = "Hello everybody"
)

func main() {

	//IN GO, ALL VARIABLES NEED TO BE USED

	fmt.Println("hello world - Noah Lambaria :)")
	//One way to initialize a variable (if uninitialized use this most likely)
	//var test int = 3
	//constants:
	//const a int = 14

	//Other way to initialize a variable (simpler)
	i := 400

	n := 1 == 1

	if n {
		fmt.Println("true")
	}

	fmt.Println(test)
	fmt.Println(i)

	//**********************************************
	//Arrays:

	numbers := [...]int{0, 1, 2, 3}

	fmt.Println(numbers)

}

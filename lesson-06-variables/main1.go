/*
1
Go Variable Types
In Go, there are different types of variables, for example:

	int- stores integers (whole numbers), such as 123 or -123
	float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
	string - stores text, such as "Hello World". String values are surrounded by double quotes
	bool- stores values with two states: true or false
*/

/*
2
In Go, there are two ways to declare a variable:
	1
	var variablename type = value

	2 (variable is inferred(compiler decides the type of the variable based on the value))
	variablename := value

	!!!Note: It is not possible to declare a variable using :=, without assigning a value to it.

*/

/*
3-types

string
int
float32
bool


*/

package main

import (
	"fmt"
)

func main1() {
	//variables
	var name string = "vikash" //String types variable (declaring type1)
	var age = 23               //inferred typed variable (declaring type1)
	lName := "Fernando"        //inferred types variable (declaring type2)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(lName)
	fmt.Println("********************")

	//Variable Declaration Without Initial Value
	var weight int
	fmt.Println(weight)

	//Value Assignment After Declaration
	var mName string
	mName = "deshitha"
	fmt.Println(mName)

}

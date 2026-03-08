//it is possible to declare multiple variables on the same line

package main

import (
	"fmt"
)

func main3() {

	var maths, science, bio int = 71, 89, 69 //type declared
	var a, b, c = "vikash", 21, "Fernando"   //without type declared (If the type keyword is not specified, you can declare different types of variables on the same line)

	fmt.Println(maths)
	fmt.Println(science)
	fmt.Println(bio)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("Multiple variable declarations can also be grouped together into a block for greater readability")
	//Multiple variable declarations can also be grouped together into a block for greater readability

	var (
		e int
		f int    = 1
		g string = "hello"
	)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

}

/*
Go Variable Naming Rules

A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.)

Go variable naming rules:
	A variable name must start with a letter or an underscore character (_)
	A variable name cannot start with a digit
	A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
	Variable names are case-sensitive (age, Age and AGE are three different variables)
	There is no limit on the length of the variable name
	A variable name cannot contain spaces
	The variable name cannot be any Go keywords
	Multi-Word Variable Names
	Variable names with more than one word can be difficult to read.


There are several techniques you can use to make them more readable:
	1 Camel Case (Each word, except the first, starts with a capital letter)
		myVariableName = "John"

	2 Pascal Case (Each word starts with a capital letter)
		MyVariableName = "John"

	3 Snake Case (Each word is separated by an underscore character)
		my_variable_name = "John"

*/

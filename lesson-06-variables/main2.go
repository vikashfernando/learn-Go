/*
Difference Between var and :=

var
 -Can be used inside and outside of functions
 -Variable declaration and value assignment can be done separately

:=
 -Can only be used inside functions
 -Variable declaration and value assignment cannot be done separately (must be done in the same line)


*/

package main

import (
	"fmt"
)

var a int
var b string = "colombo"

func main2() {
	a = 2
	fmt.Println(a)
	fmt.Println(b)

	duration := 321
	fmt.Println(duration)
}

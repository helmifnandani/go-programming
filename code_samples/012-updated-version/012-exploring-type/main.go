package main

import (
	"fmt"
)

var y = 42

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"

var z string = "Shaken, not stirred"

var a string = `James said, 
"Shaken, 

not stirred"`

// '' ambil semuanya yang ada di dalam '' literally, termasuk space, enter dan ""

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a) //"%T\n" untuk tau type dari "a"
	// z = 43 ngga boleh kayak gini
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}

package main

import "fmt"

const greet = "hello"

//var can also be used here
//:= shorthand does not work outside fxn

func main() {
	const name string = "golang" //this wont change later
	const age = 34

	//constant grouping - declaring multiple const together
	const (
		var1 = "hi"
		var2 = "welcome"
	)
	fmt.Println(var1,var2)
}
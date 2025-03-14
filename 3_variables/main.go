package main

import "fmt"

func main() {
	var name string = "ram" //cant leave a variable unused, it gives error
	fmt.Println(name)
	var fruit = "mango" //auto infer
	fmt.Println(fruit)
	var truth = false
	fmt.Println(truth)
	var age int = 34
	fmt.Println(age)

	language := "golang" //shorthand
	fmt.Println(language)
	//need a type if we are assigning a value later
	//auto infer cant be used there
	var v1, v2 int = 1, 2
	v3, v4 := 1, 2
	v5:=foo() //the datatype is not clear
	//var v5 string = foo() good practice
	const myConst string = "hello" //empty const cant be declared
	//we need to provide it a value
	
}

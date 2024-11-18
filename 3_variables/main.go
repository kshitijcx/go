package main

import "fmt"

func main() {
	var name string = "ram" //cant leave a variable unused, it gives error
	fmt.Println(name)
	var fruit = "mango" //auto infer
	fmt.Println(fruit)
	var truth = false
	fmt.Println(truth)
	var age int =34
	fmt.Println(age)

	language:="golang" //shorthand
	fmt.Println(language)
	//need a type if we are assigning a value later
	//auto infer cant be used there
	
}

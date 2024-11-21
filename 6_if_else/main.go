package main

import "fmt"

func main() {
	age := 18

	// if age >= 18 {
	// 	fmt.Println("adult")
	// }else{
	// 	fmt.Println("not adult")
	// }

	if age>=18{
		fmt.Println("adult")
	}else if age>=12{
		fmt.Println("teenager")
	}else{
		fmt.Println("kid")
	}

	var role = "admin"
	var hasPermission = true

	if role=="admin" || hasPermission {
		fmt.Println("resource accessed")
	}

	//we can declare variable inside if 
	if newAge:=3;newAge>=18{
		fmt.Println("adult")
	}else{
		fmt.Println("kid")
	}

	//there is no ternary operator in go till now, use if else 

}
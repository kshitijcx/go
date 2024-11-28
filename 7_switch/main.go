package main

import "fmt"
import "time"

func main() {
	i := 5
	//simple switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	//type switch
	whoAmI := func(i interface{}) {
		// switch t := i.(type) 
		switch i.(type){
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("other")
		}
	}
	whoAmI(45)
}

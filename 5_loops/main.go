package main

import "fmt"

func main() {
	//for is the only construct in go for looping
	i := 1
	for i <= 5 { //while loop
		fmt.Println(i)
		i++
	}

	for j:=1; j<=10; j++{ //for loop
		fmt.Println(j)
	}

	for k:=5;k<=10;k++{ //continue and break
		if k==7 {
			continue
		}
		fmt.Println(k)
	}

	for i:=range 3 {
		fmt.Println(i)
	}

	// for{ //infinite loop
	// 	fmt.Println("hello")
	// }
}

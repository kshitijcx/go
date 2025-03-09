package main

import "fmt"

func main() {
	//arrays properties
	//fixed length
	//same type
	//indexable
	//contiguous in memory
	var intArr [3]int32
	intArr[2] = 123
	fmt.Println(intArr[0:3])

	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	intArr3 := [3]int32{3, 4, 5}
	intArr4 := [...]int32{3, 4, 5}

	fmt.Println(intArr2, intArr3, intArr4)

	//slices are wrappers above arrays
	var intSlice []int32 = []int32{2, 4, 5}
	fmt.Println(len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 67) //if full doubles the space then adds the number
	fmt.Println(len(intSlice), cap(intSlice))
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) //spread operator

	// var intSlice3 = make([]int32, 3, 6) //datatype, size, capacity of slice

	//map key:value pairs
	var myMap = map[string]uint8{"ram": 32, "sam": 75}
	fmt.Println(myMap["sam"])  //if value is not avalaible, then return default value of that datatype (0 here)
	var age, ok = myMap["ram"] //maps return 2 values, 2nd one is boolean if value is found
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("not found")
	}
	delete(myMap, "ram") //delete value

	//iterating
	for name, age := range myMap {
		fmt.Println(name, age)
	}
	for index, num := range intArr {
		fmt.Println(index, num)
	}
}

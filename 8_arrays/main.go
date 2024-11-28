package main

import "fmt"

func main() {
	var nums [4]int
	//length
	fmt.Println(len(nums))

	fmt.Println(nums) //default 0

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	fmt.Println(nums[0])
	fmt.Println(nums)

	var vals [2]bool
	fmt.Println(vals) //default false

	var names [3]string
	fmt.Println(names) //default empty array

	number:=[3]int{1,2,3}
	fmt.Println(number)

	twoD:=[2][2]int{{1,2},{3,4}}
	fmt.Println(twoD)
	//fixed size, memory optimize, constant time access
	//most of the time, slices are used in golang
}

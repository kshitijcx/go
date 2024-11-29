package main

import (
	"fmt"
	"slices"
)

//slices are dynamic arrays that dont need a size
func main() {
	var nums []int //uninitialized slice is nil
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums)) //0 length

	var numbers = make([]int, 2) //specify a initial size 
	fmt.Println(numbers) //default 0 for int
	fmt.Println(cap(numbers)) //capacity 

	var numbers2 = make([]int,2,5) //size,capacity
	// var numbers2 = make([]int,0,5) //to prevent initial zeroes
	numbers2 = append(numbers2,1) //add to last
	numbers2 = append(numbers2,2) //add to last
	numbers2 = append(numbers2,3) //add to last
	numbers2 = append(numbers2,4) //add to last
	fmt.Println(numbers2)

	numbers3:=[]int{}
	numbers3=append(numbers3, 2)
	fmt.Println(numbers3)

	//elements can be added and accessed using indices just like arrays

	copy(numbers2,numbers3) //copies min len(source) into min len(dest)  
	fmt.Println(numbers2,numbers3)

	//slice operator
	var val = []int {1,2,3}
	fmt.Println(val[0:2])

	//slices package -> used for operations on slices
	fmt.Println(slices.Equal(numbers2,numbers3))

	//2D slices
	var twoD = [][]int{{1,2},{3,4}}
	fmt.Println(twoD)
}
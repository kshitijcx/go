package main

import "fmt"

func main() {
	fmt.Println(1)                  //int
	fmt.Println("this is a string") //string
	fmt.Println(true)               //bool
	fmt.Println(3.46)               //float
	//int, int8, int16, int32, int64 -> default 32 or 64
	//utint8, uint16, uint32, uint64
	//float32, float64 -> no float datatype
	//+ - * /
	// var one int16=3
	// var two int32=4
	// var three = one + two //not allowed
	// datatype mismatch
	//casting is required one+int16(two)
	//" " or ` ` for string // " " for single line
	//len gives number of bytes not length of strings
	//import "unicode/utf8"
	//utf8.RuneCountInString("string") //to get length of string
	var myRune rune = 'a' 
	fmt.Println(myRune) //rune is another datatype
	//default value for numnerical datatypes and runes is 0
	//and for string is "", bool is false
}

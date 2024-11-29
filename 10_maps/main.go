package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	m["name"] = "kshitij"
	m["hero"] = "superman"
	fmt.Println(m["name"])
	fmt.Println(m["hero"])
	fmt.Println(m["class"]) //if no key then returns empty value, 0 for int, false for bool
	fmt.Println(len(m))
	delete(m, "hero") //pass map and the key
	fmt.Println(m)
	clear(m) //empty a map
	fmt.Println(m)

	mp := map[string]int{"age": 30,"phone":1232}
	mp2 := map[string]int{"age": 30,"phone":132}
	fmt.Println(mp)

	val,ok:=mp["phone"] //returns value stored and bool

	if ok {
		fmt.Println(val)
	}else {
		fmt.Println("no val")
	}

	fmt.Println(maps.Equal(mp,mp2)) //key value datatype should be same for comparison 
}

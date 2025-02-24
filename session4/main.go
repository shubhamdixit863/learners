package main

import "fmt"

func Comoparison() {
	// maps and slices are reference types
	// m1 := make(map[string]int)
	// m1["hello"] = 1
	// m2 := make(map[string]int)
	// //fmt.Println(m1 == m2) // you cannot compare two maps ,you can compare them only with nil
	// slices cannot be compared too
	// s1 := make([]int, 6)
	// s2 := make([]int, 6)
	// fmt.Println(s1 == s2)

	//Arrays can be compared
	a1 := [2]int{1, 2}
	a2 := [2]int{1, 2}
	fmt.Println(a1 == a2)

}

func main() {
	// p := make(map[string]int)
	// p["name"] = 2
	// p["john"] = 3 // setting
	// fmt.Println(p)
	// fmt.Println(p["john"]) //getting

	//var m map[string]int // nil, no storage
	// fmt.Println(m == nil)
	// fmt.Println(m["hello"])
	//m["hello"] = 2

	//fmt.Println(m1 == m2) // you cannot compare two maps ,you can compare them only with nil
	Comoparison()
}

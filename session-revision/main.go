package main

import "fmt"

const h = 90

var c int

func main() {

	// fmt.Println("Hello  world")

	// // first way
	// //var c int
	// // var c float64
	// // var c bool
	// // c = 9
	// // fmt.Println(c)

	// //h := 9
	// fmt.Println(c)
	// add.Sum()
	// something.KL()
	// add.KK()

	var arr [6]int
	fmt.Println(arr)

	//var slc []int

	//fmt.Println(slc == nil)

	// slc2 := make([]int, 2, 4)
	// //fmt.Println(slc2)
	// fmt.Println(len(slc2), cap(slc2))
	// //fmt.Println(slc2)

	// slc2 = append(slc2, 9)
	// fmt.Println("Address of locker room", &slc2[0])
	// slc2 = append(slc2, 19)
	// fmt.Println("Address of locker room", &slc2[0])

	// slc2 = append(slc2, 190)
	// fmt.Println("Address of locker room", &slc2[0])
	// fmt.Println(slc2)
	// fmt.Println(len(slc2), cap(slc2))

	// mp := make(map[string]string)
	// mp["name"] = "shubham"
	// fmt.Println(mp)
	// fmt.Println(mp["name"])

	c := 9

	d := &c

	fmt.Println(d)

	fmt.Println(*d) // de referencing
	foo(d)          // arguments
	//foo(c)
	//a, _ := foo2(d) /// you can ignore values using _
	variadicFunction(9, 9, 9, 9, 9, 9)
}

func foo(d *int) int {
	return 9
}

func foo2(d *int) (int, string) {
	return 9, "nine"
}

func variadicFunction(n ...int) {
	fmt.Println(n)
}

func variadicFunction2(u string, n ...int) {
	fmt.Println(n)
}

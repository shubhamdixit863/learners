package main

import (
	"fmt"
	"session2/services"
)

// functions
func main() {
	// c := square(8)
	// fmt.Println(c)
	// d := sumOfTwoValues(9, 9)
	// fmt.Println(d)
	c := services.GetData()
	fmt.Println(c, services.DATA)

	a, b := Divide(8, 2)
	fmt.Println(a, b)

	a1, err := Divide(8, 0)
	// you have to put the error check after its being called
	// handling error gracefully
	if err != nil {
		fmt.Println(err)
		//	return

	}
	fmt.Println(a1)

	// ignoring the second variable

	a4, _ := Divide(8, 0)
	// you have to put the error check after its being called
	// handling error gracefully

	fmt.Println(a4)

	variadicFun(6.8, 2, 7, 9, 90, 78, 90)

	fmt.Printf("I got 90%% \n")

}

func square(num int) int {
	// you will perform the operation
	c := num ^ 2
	return c
}

func sumOfTwoValues(num1 int, num2 int) int {
	// you will perform the operation
	c := num1 + num2
	return c
}

// multiple return values
func Divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("You cannot divide by zero %s", "divide by zero")
	}

	g := num1 / num2

	return g, nil

}

// variadic function

// a function in which you can pass n number of argument

func variadicFun(f float64, n ...int) { // variadic params need to be put in th end
	fmt.Println(f, n)

	// // we will see about for loop too
	// for i := 0; i < len(n); i++ {
	// 	fmt.Println(n[i])
	// }

	//return

}

// this is easy to understand
func namedReturn(n int) (double int) {
	double = 2 * n
	return

}

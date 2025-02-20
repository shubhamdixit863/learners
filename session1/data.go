package main

import "fmt"

/*
*
In Go, when you use fmt.Scanf to read input, the %s format specifier reads a

	string until it encounters a whitespace character (space, tab, or newline). However, when you press Enter after typing the input, a newline character (\n) is also left in the input buffer.

If you don’t account for the newline, the next scan

	operation may read this leftover newline instead of waiting
	 for new input. This often leads to confusing behavior where it seems like the second scan was "skipped."

*
*/
func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name) // Newline remains in the buffer after this scan

	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age) // Reads the leftover newline instead of waiting for input

	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

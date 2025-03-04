package main

import "fmt"

type Human struct {
	name     string
	age      int
	isPlaced bool
	address  string
}

// method here // value receiver
func (h Human) Sleep() {
	fmt.Println("Good night", h.name)
}

func main() {
	// object
	h := &Human{
		name:     "sanskruti",
		age:      22,
		isPlaced: true,
		address:  "Pune",
	}

	c := 9
	fmt.Println(&c)
	fmt.Println(&h.address)

	// h2 := &Human{
	// 	name: "Narendar",
	// }
	// fmt.Println("Address h2", h2.address)
	// hPpointer := new(Human) // use of new keyword
	// fmt.Println(hPpointer)
	// //hPpointer.address = "Mumbai"

	// //hPointer := &h // pointer to the object

	// h.Sleep()

}

package main

import (
	"encoding/json"
	"fmt"

	"session8/filehandling"
)

type Data struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

// value receivers donot modify the underlying object data
// func (d Data) DoubleAge() int {
// 	d.Age = 2 * d.Age
// 	return d.Age
// }

// func (d *Data) DoubleAgePointer() int {
// 	d.Age = 2 * d.Age
// 	return d.Age
// }

func (d *Data) MarshalIt() string {
	s, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}

	return string(s)
}

func main() {
	// data := Data{
	// 	Name: "John",
	// 	Age:  23,
	// }

	// // s, err := json.Marshal(data)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }

	// fmt.Println(data.MarshalIt())

	// // c := data.DoubleAge()
	// // fmt.Println(c)
	// // fmt.Println(data.age)
	// j := data.DoubleAgePointer()
	// fmt.Println(j)
	// fmt.Println(data.age) // underlying data is modified
	//s := []byte(`{"name":"shubham","age":29}`) // string literal reprsentation
	//d := &Data{}
	//
	//err := json.Unmarshal(s, d)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(d)

	file, err := filehandling.ReadFile("../data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

}

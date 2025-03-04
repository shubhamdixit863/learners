package main

import (
	"fmt"
	"session9/internal/services"
	"strings"
)

func Add(i interface{}) interface{} {
	//	num:=2*i
	//return 2 * i
	//

	// type assertions
	//num, ok := i.(int)
	//if ok {
	//	return num * 2
	//}
	//
	//// type assertion for float values
	//
	//num2, ok := i.(float64)
	//
	//if ok {
	//	return num2 * 2
	//}
	//
	//return nil

	// switch case type assertion
	switch v := i.(type) {
	case float64:
		return v

	case string:
		return strings.ToLower(v)

	default:
		return fmt.Sprintf("%v", v)

	}

}

func main() {
	var service services.Service
	service = services.NwEmployeeServiceV2()
	fmt.Println(service.GetData())
	Add(9)
	fmt.Println(Add(9))
	fmt.Println(Add(9.8))

}

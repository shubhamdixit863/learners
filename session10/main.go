package main

import (
	"errors"
	"fmt"
)

//type CustomError struct {
//	message string
//}
//
//// CustomError wont be of type error
//// &CustomERror
//
//func (c *CustomError) Error() string {
//	return "CustomError"
//
//}
//
//func Divide(a, b int) (int, error) {
//	if b == 0 {
//		//err := errors.New("Cannot divide by 0")
//
//		return 0, &CustomError{message: "Cannot divide by 0"}
//	}
//
//	return a / b, nil
//}
//
//func ReturnCustomError() error {
//	return &CustomError{}
//}
//
//func main() {
//	_, err := Divide(6, 0)
//	if err != nil {
//		log.Println(err.Error())
//	}
//
//	//b1, err := Divide(6, 0)
//	//if err != nil {
//	//	log.Println(err.Error())
//	//}
//	//log.Println(b, b1)
//
//	if errors.Is(err, &CustomError{message: "Cannot divide by 0"}) {
//		log.Println("they are same type")
//	}
//}

type NotFoundError struct {
	Item string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Item)
}

func (e *NotFoundError) Is(err error) bool {
	var nfe *NotFoundError
	ok := errors.As(err, &nfe)
	//_, ok := err.(*NotFoundError) // type assertion
	return ok
}

func FindItem(item string) error {
	return &NotFoundError{Item: "Something"}

}

//
//func Divide() error {
//	return errors.New("some error")
//}

//type Data struct {
//}
//
//func (d *Data) Write(p []byte) (n int, err error) {
//	return 0, err
//}
//func (d *Data) String() string {
//	//return fmt.Sprintf()
//}

//	func ReadFile(fileName string) {
//		_, err := os.ReadFile(fileName)
//		if err != nil {
//			panic("Error in reading file") // it will exit the program
//
//		}
//
// }
//
//	func Recover() {
//		r := recover()
//		if r != nil {
//			log.Println("Panic recovered all good,make sure you enter correct file path")
//		}
//
//		fmt.Println("Recovered Completely")
//	}
//var err = errors.New("Some eroror")

//func ReturnError() error {
//	return fmt.Errorf("%v",)
//}

//	func ReturnError2() error {
//		return io.EOF
//	}
func main() {

	//defer fmt.Println("1")
	//
	////time.Sleep(3 * time.Second)
	//defer fmt.Println("2")
	//fmt.Println("8")
	err := FindItem("book")
	if errors.Is(err, &NotFoundError{Item: "O"}) {
		fmt.Println("Book not found!")
	}
	//	err := Divide()
	//
	//	var notFoundErr *NotFoundError
	//	boolData := errors.As(err, &notFoundErr) // it converts one error to another generally
	//	fmt.Println(boolData)
	//obj := &Data{}
	//fmt.Println(obj)
	//fmt.Fprint(obj)

	//go func() {
	//	defer Recover()
	//}()

	//	ReadFile("fileName")

	// to check whether its the same error instance
	//err1 := ReturnError()
	//
	//if errors.Is(err, err1) {
	//	log.Println("They are same error object")
	//}

	//e1 := ReturnError2()
	//
	//if errors.Is(e1, io.EOF) {
	//	log.Println("They are same error")
	//}

}

// Create a custom error for different payment methods
// And return errors from methods ,if the operation fails as the second arggyment

// internal/service
// internal/utils /errors.go // Paypal error ,

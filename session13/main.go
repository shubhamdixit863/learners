package main

import (
	"log"
	"sync"
	"time"
)

func foo() {
	log.Println("Hey there buddy")
}

// Not a right away of preventing the main go routine from exiting
//func main() {
//	go func() {
//		for i := 0; i < 1000000; i++ {
//			fmt.Println(i)
//			time.Sleep(1 * time.Second)
//		}
//
//	}()
//
//	go foo()
//
//	time.Sleep(10 * time.Second) // we are making main go routine sleep for 10 seconds
//	//fmt.Println("Main routine exiting")
//	fmt.Scanln() // a bad way of doing things
//}

// Right way of doing  is using waitgroup

func main() {
	var wg sync.WaitGroup

	go func() {

		log.Println("Goroutine 1")
		wg.Done()
	}()
	go func() {
		c := 9
		log.Println("Goroutine 2", c)
		wg.Done()
	}()

	go func() {
		wg.Done()
		log.Println("Goroutine 3")
		time.Sleep(10 * time.Second)

	}()

	wg.Add(3)

	wg.Wait() // will actually hold the go routine from exiting

	log.Println("Now the wait is over exiting the main routine")

}

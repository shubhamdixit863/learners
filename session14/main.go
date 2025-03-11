package main

import (
	"fmt"
)

//var c = 9
//var wg sync.WaitGroup
//
//func a() {
//	defer wg.Done()
//	c = 8
//}
//
//func b() {
//	defer wg.Done()
//	fmt.Println("exiting", c)
//	if c >= 9 {
//		fmt.Println("hello there")
//	}
//	fmt.Println("exiting", c)
//}

func main() {
	//var wg sync.WaitGroup
	//// unbuffered channel
	//ch := make(chan int)
	//go func() {
	//	defer wg.Done()
	//	ch <- 9
	//	close(ch) // sender should close the channel
	//
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	v := <-ch
	//	fmt.Println("Data coming from chan", v)
	//}()
	//
	//wg.Add(2)
	//wg.Wait()
	//	var wg sync.WaitGroup
	slc := []int{8, 1, 3, 90, 2}
	//ch := make(chan int)
	//	chbuf := make(chan int, 5)
	//fmt.Println(slc)
	//sorting.BubbleSort(slc)
	//fmt.Println(slc)
	//
	//go sorting.BubbleOuter(slc, ch, &wg)
	//go sorting.BubbleInner(slc, ch, &wg)
	//go sorting.BubbleOuter2(slc, chbuf, &wg)
	//go sorting.BubbleInner2(slc, chbuf, &wg)
	//wg.Add(2)
	//wg.Wait()
	fmt.Println(slc)
}

package sorting

import (
	"fmt"
	"sync"
)

func BubbleSort(slc []int) {
	for i := 0; i < len(slc); i++ {
		for j := 0; j < len(slc)-1; j++ {
			// swap
			if slc[j] > slc[j+1] {
				temp := slc[j+1]
				slc[j+1] = slc[j]
				slc[j] = temp
			}
		}
	}
}

func BubbleOuter(slc []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(slc); i++ {
		ch <- i
		fmt.Println("From outer", i)
	}

	close(ch)
}

func BubbleInner(slc []int, ch chan int, wg *sync.WaitGroup) {

	//v, ok := <-ch
	//fmt.Println(ok, v)
	//if !ok {
	defer wg.Done()
	//}

	for v := range ch {
		fmt.Println("from inner", v)
		for j := 0; j < len(slc)-1; j++ {
			// swap
			if slc[j] > slc[j+1] {
				temp := slc[j+1]
				slc[j+1] = slc[j]
				slc[j] = temp
			}
		}
	}

	fmt.Println("I am exiting now bye")

}

func BubbleOuter2(slc []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(slc); i++ {
		ch <- i
		fmt.Println("From outer", i)
	}

	close(ch)
}

func BubbleInner2(slc []int, ch chan int, wg *sync.WaitGroup) {
	var ok = true
	//v, ok := <-ch
	//fmt.Println(ok, v)
	defer wg.Done()
	for ok {
		_, ok = <-ch
		for j := 0; j < len(slc)-1; j++ {
			// swap
			if slc[j] > slc[j+1] {
				temp := slc[j+1]
				slc[j+1] = slc[j]
				slc[j] = temp
			}
		}

	}

	//for v := range ch {
	//	fmt.Println("from inner", v)
	//
	//}

	fmt.Println("I am exiting now bye")

}

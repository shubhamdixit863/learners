package main

import (
	"fmt"
	"math/rand"
)

// fan in pattern along with fan out

func generator(genCh chan<- int, count int) {
	// it will generate random numbers
	for i := 0; i < count; i++ {
		rn := rand.Intn(100)
		genCh <- rn
	}

	close(genCh)
}

func worker(in chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			sq := v * v
			out <- sq
		}
		close(out)
	}()

	return out
}

func merge(resultCh chan int, n ...<-chan int) {
	for _, v := range n {
		for d := range v {
			resultCh <- d

		}
	}
	close(resultCh)
}

const NUM_WORKERS = 5
const COUNT = 1000

func main() {
	genCh := make(chan int)
	resultCh := make(chan int)
	//var sliceOfChannels []<-chan int
	sliceOfChannels := make([]<-chan int, 0)
	fmt.Println(cap(sliceOfChannels))
	go generator(genCh, COUNT)
	for range NUM_WORKERS { // Workers will be spawned
		outCh := worker(genCh)
		sliceOfChannels = append(sliceOfChannels, outCh)
	}
	go func() {
		merge(resultCh, sliceOfChannels...)
	}()

	for v := range resultCh {
		fmt.Println(v)

	}

}

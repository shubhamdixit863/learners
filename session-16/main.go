package main

import (
	"log"
	"sync"
)

var c = 9 // shared memory

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// crtical section
			log.Println(1)
			log.Println(1)

			log.Println(1)

			log.Println(1)

			log.Println(1)

			mu.Lock()
			c = c + 1
			mu.Unlock()
		}()
	}
	wg.Wait()
	log.Println(c)

}

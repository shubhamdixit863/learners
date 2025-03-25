package main

import (
	"fmt"
	"strings"
	"sync"
)

func worker(taskCh chan string, wg *sync.WaitGroup, workerId int, resultCh chan string) {
	// Any job
	defer wg.Done()
	for v := range taskCh {
		s := fmt.Sprintf("Workerid %d - %s\n", workerId, strings.ToUpper(v))
		resultCh <- s // statement is blocking
	}

}

const NUM_WORKERS = 2

func main() {
	var wg sync.WaitGroup
	var rg sync.WaitGroup
	taskChan := make(chan string)
	resultchan := make(chan string)
	fn := []string{"a.txt", "b.txt", "c.txt", "d.txt"}

	for i := 0; i < NUM_WORKERS; i++ {
		wg.Add(1)
		go worker(taskChan, &wg, i, resultchan)
	}
	rg.Add(1)
	go func() {
		defer rg.Done()
		for g := range resultchan {
			fmt.Println(g)
		}
		fmt.Println("All done")
	}()

	for i := 0; i < len(fn); i++ {
		taskChan <- fn[i]
	}

	close(taskChan)
	wg.Wait()
	close(resultchan)
	rg.Wait()

}

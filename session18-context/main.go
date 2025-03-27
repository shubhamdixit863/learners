package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	//
	wg.Add(1)
	//ctx, cancel := context.WithCancel(context.Background())
	//go func(ctx context.Context) {
	//	defer wg.Done()
	//	<-ctx.Done() // line will block
	//	log.Println("Returning from go routine")
	//}(ctx)
	//
	//cancel()
	//wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	go func(ctx context.Context) {
		defer wg.Done()
		<-ctx.Done()
		log.Println("I am quitting now as 5 seconds are over")
	}(ctx)

	cancel()
	time.Sleep(15 * time.Second)

	wg.Wait()
}

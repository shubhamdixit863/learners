package main

import (
	"sync"
	"sync/atomic"
)

func sender(ch chan int) {
	i := 0
	for {

		ch <- i
		i++
	}
}

/*
*
The line you're asking about is:

```go
ch <- 9
```

And your question is **why sending `9` to the channel is not blocking** even though the channel has a **buffer size of 2**, and you already sent:

```go
ch <- 3
ch <- 4
```

Let’s break it down and see **why** you're not seeing a deadlock here *immediately*, and *why it could still deadlock or hang silently in some cases.*

---

### 🔍 Breakdown of Code Behavior:

```go
ch := make(chan int, 2)
```

You've created a buffered channel with **capacity 2**.

Then:

```go
ch <- 3 // 1st value
ch <- 4 // 2nd value (buffer now full)
log.Println("Sent to channel")
ch <- 9 // 3rd value, this blocks unless someone reads from the channel
```

The key point is:
> **ch <- 9** will block if the buffer is full and there's no goroutine currently reading from the channel.

---

### 🧠 So why is it NOT blocking in your case?

Let's look at this part of the code:

```go

	go func() {
		for {

		}
		fmt.Println(<-ch)
		wg.Done()
	}()

```

**This goroutine will never read from the channel.**
Why?

Because `for {}` is an **infinite empty loop** — the code below it (`fmt.Println(<-ch)`) is **never reached**.

That means:
- The receiver goroutine is doing nothing.
- No one is consuming from the channel.
- **Buffer size is 2**, so after sending 3 and 4, it's full.
- `ch <- 9` **should block** — BUT...

---

### 🤯 So, why does it *seem* to not block?

Actually, **it *does* block**, and that goroutine is stuck there. It never completes and calls `wg.Done()`. You're just **not seeing** the deadlock message because:

- **Go only throws a fatal `deadlock` panic if all goroutines are asleep** and there's nothing progressing.
- In your case, that sender goroutine is blocked on `ch <- 9`, and another goroutine is running an infinite loop `for {}` — so Go thinks the program is still "doing something".
- As a result, **you never see a panic**, but the program is in a **soft deadlock / hang**.

---

### ✅ To confirm the deadlock

Change the infinite loop to a `time.Sleep` or remove it:

```go

	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()

```

Now `ch <- 3` will be read, freeing space in the buffer. Then `ch <- 4` and `ch <- 9` will go through without blocking.

Or just add a `log.Println("stuck")` before `ch <- 9` to see that it's never moving past that point:

```go
log.Println("Sending 9...")
ch <- 9
log.Println("Sent 9") // This won't print
```

---

### 🧼 Final Suggestion

If you want to **avoid blocking** and see the behavior clearly, try this:

```go

	go func() {
		for val := range ch {
			fmt.Println("Received:", val)
		}
	}()

```

Let it continuously receive, and the buffer will never overflow unless the producer outpaces the consumer.

---

### ✅ Summary

- `ch <- 9` **does block**, because the buffer is full and no one is reading.
- The deadlock is *hidden* because your other goroutine is running an infinite loop (`for {}`), keeping the scheduler alive.
- Replace the `for {}` with a proper channel reader to observe the correct behavior.

Let me know if you want help writing a version that shows this clearly!
*/
//func main() {
//	//var wg sync.WaitGroup
//	//ch := make(chan int, 2)
//	//wg.Add(2)
//	//go func() {
//	//	ch <- 3
//	//	ch <- 4
//	//	log.Println("Sent to channel")
//	//	ch <- 9
//	//	wg.Done()
//	//}()
//
//	//for i := 0; i < 10; i++ {
//	//	wg.Add(1)
//	//	go func(id int) {
//	//		defer wg.Done()
//	//		for v := range ch {
//	//			log.Printf("Go routine Id -%d ,value - %d", i, v)
//	//		}
//	//	}(i)
//	//}
//	//
//	//go sender(ch)
//	//wg.Wait()
//	//go func() {
//	//	for {
//	//
//	//	}
//	//	fmt.Println(<-ch)
//	//	wg.Done()
//	//}()
//
//	//wg.Wait()
//}

var counter int64 = 0

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	//mu.Lock()
	//counter = counter + 1
	//mu.Unlock()
	atomic.AddInt64(&counter, 1)

}

//func main() {
//	//var wg sync.WaitGroup
//	ch1 := make(chan int)
//	ch2 := make(chan string)
//	//var mu sync.Mutex
//	////wg = &sync.WaitGroup{}
//	//for i := 0; i < 10; i++ {
//	//	wg.Add(1)
//	//	go increment(&wg, &mu)
//	//}
//	//
//	//wg.Wait()
//
//	go func() {
//		time.Sleep(10 * time.Second)
//		ch1 <- 9
//		//close(ch1)
//	}()
//
//	go func() {
//		time.Sleep(10 * time.Second)
//		ch2 <- "hello world"
//
//		//close(ch2)
//	}()
//
//	for {
//
//		select {
//		case v := <-ch1:
//			// You can do anything with value
//			fmt.Println(v)
//		case g := <-ch2:
//			fmt.Println(g)
//
//			default:
//				fmt.Println("No channel ready")
//			}
//
//			time.Sleep(2 * time.Second)
//
//		}
//	}
//}

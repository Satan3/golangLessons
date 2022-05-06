package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//AddMutex()
	//AddAtomic()

	//StoreLoadSwap()
	//compareAndSwap()

	atomicVal()
}

func AddMutex() {
	start := time.Now()

	var (
		counter int64
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("With mutex:", time.Now().Sub(start).Seconds())
}

func AddAtomic() {
	start := time.Now()
	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("With atomic:", time.Now().Sub(start).Seconds())
}

func StoreLoadSwap() {
	var counter int64

	fmt.Println(atomic.LoadInt64(&counter))

	atomic.StoreInt64(&counter, 5)
	fmt.Println(atomic.LoadInt64(&counter))

	fmt.Println(atomic.SwapInt64(&counter, 10))
	fmt.Println(atomic.LoadInt64(&counter))
}

func compareAndSwap() {
	var (
		counter int64
		wg      sync.WaitGroup
	)
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()

			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return
			}

			fmt.Println("Swapped goroutine number is", i)
		}(i)
	}

	wg.Wait()
	fmt.Println(counter)
}

func atomicVal() {
	var (
		value atomic.Value
	)

	value.Store(1)
	fmt.Println(value.Load())

	fmt.Println(value.Swap(2))
	fmt.Println(value.Load())

	fmt.Println(value.CompareAndSwap(2, 3))
	fmt.Println(value.Load())
}

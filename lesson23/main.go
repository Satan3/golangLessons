package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//baseKnowledge()
	workerPool()
}

func baseKnowledge() {
	ctx := context.Background()
	fmt.Println(ctx)

	toDo := context.TODO()
	fmt.Println(toDo)

	withValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println(withValue.Value("name"))

	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err())
	cancel()
	fmt.Println(withCancel.Err())

	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	fmt.Println(withDeadline.Deadline())
	fmt.Println(withDeadline.Err())
	fmt.Println(<-withDeadline.Done())

	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	fmt.Println(withTimeout.Done())
}

func workerPool() {

	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()

	wg := &sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			/*if i == 500 {
				cancel()
			}*/
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}

	fmt.Println(counter)
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}

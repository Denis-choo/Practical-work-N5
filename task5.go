package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-done:
					fmt.Printf("Воркер %d завершает запрос\n", w)
					time.Sleep(time.Second)
					fmt.Printf("Воркер %d вышел\n", w)
					return
				default:
					fmt.Printf("Воркер %d работает\n", w)
					time.Sleep(time.Second)
				}
			}
		}()
	}
	time.Sleep(2 * time.Second)
	close(done)
	wg.Wait()
}

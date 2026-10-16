package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)
	ch2 <- "обычная"
	ch2 <- "обычная"
	ch1 <- "срочная"
	ch1 <- "срочная"
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 1; i <= 4; i++ {
			select {
			case x := <-ch1:
				fmt.Println(x)
			default:
				x := <-ch2
				fmt.Println(x)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Всё!")
}

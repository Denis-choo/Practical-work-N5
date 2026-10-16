package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	tick := time.Tick(1 * time.Second)

	done := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		done <- struct{}{}
	}()

	fmt.Println("Начинаю отсчет")

	for countdown := 1; countdown > 0; countdown++ {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-done:
			fmt.Println("Запуск отменен!")
			return
		}
	}

}

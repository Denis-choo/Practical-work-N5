package main

import(
"fmt"
"time"
)

func main() {
ch :=make(chan string)

go func() {
time.Sleep(3 * time.Second)
ch <-"ответ"
}()
select {
case x:= <-ch:
 fmt.Println(x)
case <- time.After(2 * time.Second):
fmt.Println("таймаут")
	}


}
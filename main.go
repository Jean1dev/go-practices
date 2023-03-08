package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("teste", i)
		time.Sleep(1000)
	}
}

func main() {
	canal := make(chan string)

	go func() {
		canal <- "teste comunicacao thread"
	}()

	msg := <-canal
	fmt.Println(msg)
}

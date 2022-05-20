package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func producter(name string, ch chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println(name, "product :", i)
		ch <- i
	}
	wg.Done()
}

func consumer(name string, ch chan int) {

	for i := range ch {
		fmt.Println(name, "consum :", i)
	}
	fmt.Println("consumer is Done")
	wg.Done()
}

func main() {

	ch := make(chan int)

	wg.Add(2)
	go producter("生产者A", ch)
	go producter("生产者B", ch)
	wg.Wait()
	wg.Add(2)
	go consumer("消费者A", ch)
	go consumer("消费者B", ch)
	wg.Wait()
	close(ch)
}

package main

import (
	"fmt"
)

func numberGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start + i
	}
	close(out)
	// 이 코드는 여기서 close() 하지 않으면 deadlock 됨
	// A receive from an unbuffered channel happens before the send on that channel completes.
}

func printNumbers(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	// close(in) // receive-only 채널은 close를 할 수 없다.
	done <- true
}

func main() {
	numberChan := make(chan int)
	done := make(chan bool)
	go numberGen(1, 10, numberChan)
	go printNumbers(numberChan, done)

	<-done
}

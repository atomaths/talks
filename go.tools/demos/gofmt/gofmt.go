package main

import (
	"fmt"
	"log"
)

func sqrt(i int) int {
	return i * i
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(sqrt(i))
	}

	log.Print("Hello" + "World")
}

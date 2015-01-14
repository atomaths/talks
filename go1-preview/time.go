package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	timestamp := now.Unix()
	fmt.Printf("%v\n", timestamp)

	// Month()는 Month 타입 (int)가 리턴되는데 String() 메서드가 있어서 January로 프린트된다.
	fmt.Println(now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
}

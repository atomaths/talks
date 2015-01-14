package main

import (
	"fmt"
)

type Day struct {
	name string
	date int
}

func main() {
	Christmas := Day{"Christmas", 25}
	Christmas2 := Day{"Christmas", 25}
	Thanksgiving := Day{"Thanksgiving", 31}

	// Go 1 이전에는 
	// "invalid operation: Christmas == Thanksgiving (operator == not defined on struct)" 컴파일 에러 
	if Christmas == Thanksgiving {
		println("Christmas and Thanksgiving is not equal")
	}

	if Christmas == Christmas2 {
		println("Christmas and Thanksgiving is equal")
	}


	holiday := map[Day]bool{
		Christmas: true,
		Thanksgiving: true,
	}

	fmt.Printf("Christmas is a holiday: %t\n", holiday[Christmas])
}

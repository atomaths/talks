package main

import (
	"fmt"
	"strconv"
)

type MyError struct {
	MyField int
}

func (m *MyError) Error() string {
	return "MyError implements the error interface. MyField is " + strconv.Itoa(m.MyField)
}

func (m *MyError) Sum(x, y int) (n int, err error) {
	return x+y, m
}



func main() {
	myerr := new(MyError)
	myerr.MyField = 200

	n, err := myerr.Sum(1, 2)

	fmt.Println(n, err)
}

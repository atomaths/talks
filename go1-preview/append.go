package main

import (
	"fmt"
)

func main() {
	// Before Go 1
	fmt.Printf("%s\n", append([]byte("Hello "), []byte("World")...)) // Go 1에서 에러는 아님

	// After Go 1
	fmt.Printf("%s\n", append([]byte("Hello "), " World"...))
	fmt.Printf("%s\n", append([]string{"Hello "}, " World", " other strings"))

	var s Stack
	s.Push("hell") // 이렇게 사용 가능
}



type Stack struct {
	data []interface{}
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() interface{} {
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil // to avoid memory leak
	s.data = s.data[:i]
	return res
}

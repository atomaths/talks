package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Type assertions
//	x.(T)
//	x는 인터페이스의 타입. x가 nil이 아니고, x에 저장된 값은 T 타입이라고 assertion
//	만일 T가 인터페이스 타입이 아니라면 x.(T)는 x는 T와 같은 타입이라고 assertion
//	T가 인터페이스 타입이라면 x는 인터페이스 T를 구현한 것이라고 assertion

//	var x interface{}
//	x = 10
//	c := x.(int)
//	c, ok := x.(int)     // c == 10, ok == true
//	c, ok := x.(string)  // c == zeroed value, ok == false

//	인터페이스 변수 x가 MyInterface를 구현했다면 x.(MyInterface).Myfunc() 처럼 method chaining으로 사용 가능

	// Conversions
	b := []byte{'H', 'e', 'l'}
	fmt.Printf("%s\n", string(b))

	n := 10
	fn := float64(n)
	fmt.Printf("%f, %f, %s\n", n, fn, reflect.TypeOf(fn))
}

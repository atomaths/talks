package main

import (
	"fmt"
	"reflect"
)

type Date struct {
	month string
	day   int
}

func main() {
	// Struct values, fully qualified; always legal.
	holiday1 := []Date{
		Date{"Feb", 14},
		Date{"Nov", 11},
		Date{"Dec", 25},
	}
	// Struct values, type name elided; always legal.
	holiday2 := []Date{
		{"Feb", 14},
		{"Nov", 11},
		{"Dec", 25},
	}
	// Pointers, fully qualified, always legal.
	holiday3 := []*Date{
		// []*Date 표현은 pointer Date 타입들의 slice를 가리키므로
		// Date{"Feb", 14}와 같이 type *Date가 와야 할 자리에 type Date가 오면 컴파일 에러가 난다
		&Date{"Feb", 14},
		&Date{"Nov", 11},
		&Date{"Dec", 25},
	}
	// Pointers, type name elided; legal in Go 1.
	// Go 1 이전에는 “missing type in composite literal” 에러남.
	holiday4 := []*Date{
		{"Feb", 14},
		{"Nov", 11},
		{"Dec", 25},
	}

	fmt.Println(holiday1)
	fmt.Println(holiday2)
	fmt.Printf("%v\n", reflect.ValueOf(holiday3))
	fmt.Printf("%v\n", holiday4)
}

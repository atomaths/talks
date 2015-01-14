package main

import (
	"fmt"
)

func main() {
	fmt.Println(Bug())
	Foo()
	fmt.Println(Bar())
}

// Go 1 이전에는 5, 0, 100이 리턴되나
// Go 1부터는 컴파일 시, "j is shadowed during return" 에러가 난다.
func Bug() (i, j, k int) {
	for i = 0; i < 5; i++ {
		for j := 0; j < 5; j++ { // Redeclares j.
			k += i*j
			if k > 100 {
				return // Rejected: j is shadowed here.
			}
		}
	}
	return // OK: j is not shadowed here.
}

func Foo() {
	i := 10
	for j := 0; j < 5; j++ {
		i := 3
		fmt.Println("other i", i)
		// i = 3 // same i
	}
	fmt.Println(i)
}

func Bar() (i int) {
	// i := 10 // 이렇게 block scope 내가 아닌 곳에서 redeclare하면 r60.3, Go 1 모두 컴파일 에러: no new variables on left side of :=

	for i := 0; i < 5; i++ {
		if i == 4 {
			return
			// return i
		}
	}
	return
}

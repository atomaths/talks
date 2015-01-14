package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{"Jan": 31, "Feb":28, "Mar":31, "Apr":40, "May":31}

	//m["Jan"] = -1, false // r60.3 방식
	//delete(m, "Jan") // Go 1 방식

	fmt.Println(m)

	for k, v := range m {
		// This loop should not assume "Jan" will be visited first.
		fmt.Println(k, v)
	}
}

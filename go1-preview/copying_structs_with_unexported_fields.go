package main

import "fmt"
import "my_new_struct_pkg"

// Go 1 이전에는 "implicit assignment of unexported field 'secret' of my_new_struct_pkg.Struct in assignment"
// 에러가 난다
func main() {
	myStruct := my_new_struct_pkg.NewStruct(23)
	copyOfMyStruct := myStruct
	fmt.Println(myStruct, copyOfMyStruct)
}

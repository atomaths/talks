package main

func foo() {
	println("foo invoked")
}

func main() {
	f1 := foo
	f2 := foo

	// Go 1에서는 function과 map은 직접 == 비교가 안 된다.
	// invalid operation: f1 == f2 (func can only be compared to nil) 컴파일 시 이런 에러남
	// Go 1 전에는 아래 비교가 가능했다. (결과도 성공)

	if f1 == f2 {
		println("equal")
	} else {
		println("not equal")
	}
}

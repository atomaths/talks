package main

var PackageGlobal int

func init() {
    c := make(chan int)
    go initializationFunction(c)
    PackageGlobal = <-c // Go 1에서는 deadlock 없이 initializationFunction goroutine이 실행되어
			// c channel로 write 할 때까지 read 동작을 기다려준다.
			// 하지만 Go 1 이전에는 init()은 goroutine이 생성만 될뿐, 실행되지는 않기때문에
			// deadlock이 된다.
//    PackageGlobal = 20 // 이렇게 하는 건 Go 1 이전에도 된다. 아무 syncronization event가 없기 때문.
}

func initializationFunction(c chan int) {
//	c <- 10
}

func main() {
	// Go 1에서는 정상적으로 10이 출력되지만,
	// Go 1 이전 버전에서는 컴파일은 에러없이 돼도, "throw: init rescheduling" 런타임 에러가 난다.
	println(PackageGlobal)
}


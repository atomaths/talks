// OOP style in Go

package main

import (
    "fmt"
)

type Pointer interface {
    draw()
}

type Point struct {
    x, y int
}

func (p *Point) draw() {
    fmt.Println(p.x, p.y)
}

func (p *Point) otherMethod() {
    fmt.Println("otherMethod")
}

func main() {
    p := Point{3, 4}
    p.draw()
    p.otherMethod()

    np := NamedPoint{Point{5, 6}, "hello"}
    np.draw()
    np.otherMethod()
}

///////////////////////////////////////////////////
type NamedPoint struct {
    Point
    name string
}

func (np *NamedPoint) draw() {
    fmt.Println(np.x, np.y, np.name)
}

package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.height
}

func (r rectangle) perimeter() float64 {
	return (2 * r.length) + (2 * r.height)
}

func measure(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {

}

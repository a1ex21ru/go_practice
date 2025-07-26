package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	Side float32
}

func (s Square) Area() float32 {
	return s.Side * s.Side
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return c.Radius * c.Radius * math.Pi
}

func PrintShapeArea(sh Shape) {
	fmt.Println(sh.Area())
}

func PrintInterface(i interface{}) {
	str, err := i.(string)
	if !err {
		fmt.Println("interface is not string")
	}
	fmt.Println(str)
}

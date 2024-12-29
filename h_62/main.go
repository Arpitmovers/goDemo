// Hands-on exercise #62 - interfaces
// ●
//  create a type SQUARE
// ○ length float64
// ○ width float64
// ● create a type CIRCLE
// ○ radius float64
// ● attach a method to each that calculates AREA and returns it
// ○ circle area= π r 2
// ■ math.Pi
// ■ math.Pow
// ○ square area = L * W
// ● create a type SHAPE that defines an interface as anything that has the AREA method
// ● create a func INFO which takes type shape and then prints the area
// ● create a value of type square
// ● create a value of type circle
// ● use func info to print the area of square
// ● use func info to print the area of circle
// https://go.dev/play/p/8BFxl2GXgcw
// curriculum item # 155-hands-on-62-interfaces

package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	width  float64
}

type CIRCLE struct {
	radius float64
}

func (s1 SQUARE) area() float64 {
	return s1.length * s1.width
}

func (c CIRCLE) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// any value of type that has area( ) method will  also BE of type shape
// so , circle , square alsos are also of  type shape
type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {

	s1 := SQUARE{
		length: 12,
		width:  12,
	}

	c1 := CIRCLE{
		radius: 12,
	}
	sqrArea := info(s1)
	circleArea := info(c1)

	fmt.Println("sqrArea", sqrArea)
	fmt.Println("circleArea", circleArea)
}

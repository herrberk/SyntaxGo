package main

import (
	"math"
	"fmt"
)

type Shape interface{
	area() float64
}

type Rectangle struct{
	height float64
	width float64
}

type Circle struct{
	radius float64
}
//noinspection GoDuplicateFunctionOrMethod
func (r Rectangle) area() float64{
	return r.width*r.height
}
func (c Circle) area() float64{
	return math.Pi *math.Pow(c.radius,2)
}

func getArea(shape Shape) float64{
	return shape.area()
}

func main(){
	rect:= Rectangle{10,20}

	var area float64= 0;
	area = getArea(rect)
	fmt.Println(area)
}
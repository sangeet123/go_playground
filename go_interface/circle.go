package main

import "math"
import "fmt"

type shape interface{
	area() float64
	perimeter() float64
}

type circle struct{
	radius float64
}

type rectangle struct{
	length float64
	bredth float64
}


//ensure circle implements the interface
var _ shape = (*circle)(nil)

//using pass by value so that property of circle cannot be changed
func(c circle)area()float64{
	return math.Pow(c.radius, 2)*math.Pi 
}

func(c circle)perimeter()float64{
	return 2 * math.Pi * c.radius
}

//ensure rectangle implements the interface
var _ shape = (*rectangle)(nil)

func(r rectangle)area()float64{
	return r.length * r.bredth
}

func(r rectangle)perimeter()float64{
	return 2* (r.length + r.bredth)
}



func main(){
	c := circle{radius: 4}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())

	fmt.Println("-------------------------------------------------------------------------------")
	r := rectangle{length:2, bredth:4}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
}
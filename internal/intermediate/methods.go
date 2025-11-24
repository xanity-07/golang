package intermediate

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

type Shape struct {
	Rectangle
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64) *Rectangle {
	r.length *= factor
	r.width *= factor
	return r
}

func IntroMethods() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}

	area := rect.Area()
	fmt.Println("Area of rect:", area)
	biggerRect := rect.Scale(1.2)

	fmt.Println("Bigger rect:", *biggerRect)
	fmt.Println("Bigger rect Area:", biggerRect.Area())

	s := Shape{Rectangle: Rectangle{
		length: 10,
		width:  15,
	}}

	fmt.Println(s)
}

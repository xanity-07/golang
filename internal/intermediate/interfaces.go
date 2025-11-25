package intermediate

import (
	"fmt"
	"math"
)

// * Syntax
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.length + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * (math.Pi * c.radius)
}

// func (c circle) diamiter() float64 {
// 	return 2 * c.radius
// }

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func IntroInterfaces() {
	r := rect{length: 4, width: 3}
	c := circle{radius: 5}
	measure(r)
	measure(c)

	//* Accepts any val as parameters
	acceptAny(1, "string", true, map[string]string{"key": "val"})
	anyVal("")
	anyVal(4)
	anyVal(false)
}

func acceptAny(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

// * interface{} can be replaced by any
func anyVal(i interface{}) {
	switch i.(type) {
	case int:
		{
			fmt.Println("Int")
		}
	case string:
		{
			fmt.Println("Int")
		}
	case bool:
		{
			fmt.Println("Int")
		}
	default:
		{
			fmt.Println("Unkown")
		}
	}
}

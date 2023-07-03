package bridge

import "fmt"

type Shape interface {
	Area() float32
	Draw()
}

type Cycle struct {
	Radius float32
	Color  Color
}

func (c Cycle) Area() float32 {
	return 22 * c.Radius * c.Radius / 7
}

func (c Cycle) Diameter() float32 {
	return 2 * c.Radius
}

func (c Cycle) Draw() {
	fmt.Printf("Draw a cycle with diameter %.3f with color %s\n", c.Diameter(), c.Color.Hex())
}

type Rectangle struct {
	SideLength float32
	Color      Color
}

func (r Rectangle) Area() float32 {
	return r.SideLength * r.SideLength
}

func (r Rectangle) Draw() {
	fmt.Printf("Draw a rectangle with color %s\n", r.Color.Hex())
}

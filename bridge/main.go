// The objective is to reduce complexibility when having types that have the same methods
// Using an interface that applies the common methods, and for each type we should specify the method

package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	raster := RasterRenderer{}
	circle1 := NewCircle(&raster, 2)

	circle1.Draw()

	vector := VectorRenderer{}
	circle2 := NewCircle(&vector, 5)

	circle2.Draw()
}

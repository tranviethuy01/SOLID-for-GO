package main

import "fmt"

// Shape interface with an Area method
type Shape interface {
	Area() float64
}

// Square type implementing the Shape interface
type Square struct {
	SideLength float64
}

func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

// Rectangle type implementing the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function demonstrating LSP
func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	// Original shapes
	square := Square{SideLength: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}

	// Calculate area without modifying CalculateArea function
	areaSquare := CalculateArea(square)
	areaRectangle := CalculateArea(rectangle)

	fmt.Printf("Area of Square: %.2f\n", areaSquare)
	fmt.Printf("Area of Rectangle: %.2f\n", areaRectangle)

	// Introduce a new shape without modifying CalculateArea function
	circle := Circle{Radius: 3.0}
	areaCircle := CalculateArea(circle)
	fmt.Printf("Area of Circle: %.2f\n", areaCircle)
}

// Circle type implementing the Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

/*
Result


Area of Square: 25.00
Area of Rectangle: 24.00
Area of Circle: 28.26


*/


/*

Explain this example:

We have the Shape interface with an Area method.
The Square, Rectangle, and Circle types implement the Shape interface, providing their specific implementations for calculating the area.
The CalculateArea function takes any type that satisfies the Shape interface and calculates the area without knowing the specific types of shapes.
With this design, you can easily introduce new types of shapes (e.g., Circle) without modifying the existing CalculateArea function. This demonstrates adherence to the Liskov Substitution Principle (LSP) because different types of shapes can be treated uniformly in functions like CalculateArea without altering the behavior of the function.

*/

package main

import "fmt"

// Square type with a function for calculating the area
type Square struct {
	SideLength float64
}

func CalculateSquareArea(square Square) float64 {
	return square.SideLength * square.SideLength
}

// Rectangle type with a function for calculating the area
type Rectangle struct {
	Width  float64
	Height float64
}

func CalculateRectangleArea(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// Circle type with a function for calculating the area
type Circle struct {
	Radius float64
}

func CalculateCircleArea(circle Circle) float64 {
	return 3.14 * circle.Radius * circle.Radius
}

func main() {
	// Original shapes
	square := Square{SideLength: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}
	circle := Circle{Radius: 3.0}

	// Calculate area using specific functions for each shape
	areaSquare := CalculateSquareArea(square)
	areaRectangle := CalculateRectangleArea(rectangle)
	areaCircle := CalculateCircleArea(circle)

	fmt.Printf("Area of Square: %.2f\n", areaSquare)
	fmt.Printf("Area of Rectangle: %.2f\n", areaRectangle)
	fmt.Printf("Area of Circle: %.2f\n", areaCircle)
}


/*

How bad is this example:

Each shape type (Square, Rectangle, Circle) has its own function for calculating the area (CalculateSquareArea, CalculateRectangleArea, CalculateCircleArea).
The main function calculates the area for each shape using the respective functions.
Without using an interface, you would need to have separate functions for each type, and any function that needs to work with different shapes would need to explicitly handle each type individually. This can result in less flexibility and more code duplication compared to using an interface.

*/

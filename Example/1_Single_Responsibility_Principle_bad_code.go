package main

import "fmt"

// Car type
type Car struct {
	FuelLevel float64
}

// Plane type
type Plane struct {
	FuelLevel float64
}

// Method to drive a car
func (car *Car) Drive() {
	fmt.Println("Driving the car")
}

// Method to refuel a car
func (car *Car) Refuel() {
	fmt.Println("Refueling the car")
}

// Method to fly a plane
func (plane *Plane) Drive() {
	fmt.Println("Flying the plane")
}

// Method to refuel a plane
func (plane *Plane) Refuel() {
	fmt.Println("Refueling the plane")
}

func main() {
	// Create instances of Car and Plane
	car := Car{FuelLevel: 20.0}
	plane := Plane{FuelLevel: 10000.0}

	// Drive and refuel the Car
	fmt.Println("Driving and refueling the Car:")
	car.Drive()
	car.Refuel()

	fmt.Println("\nFlying and refueling the Plane:")
	// Drive and refuel the Plane
	plane.Drive()
	plane.Refuel()
}


/*


Driving and refueling the Car:
Driving the car
Refueling the car

Flying and refueling the Plane:
Flying the plane
Refueling the plane

*/

/*

How bad is this code 

Methods Drive and Refuel are defined for both Car and Plane types.
Each method handles the specific behavior associated with the respective type.
While this approach avoids interfaces and provides a direct solution for each type, it may result in code duplication if similar behavior needs to be shared among multiple types. The use of interfaces allows you to define a common contract and reuse behavior across different types.

*/

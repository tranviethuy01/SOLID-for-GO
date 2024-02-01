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

// Function to drive a vehicle (Car or Plane)
func DriveVehicleWithoutInterface(vehicle interface{}) {
	switch vehicle.(type) {
	case Car:
		fmt.Println("Driving the car")
	case Plane:
		fmt.Println("Flying the plane")
	default:
		fmt.Println("Unknown vehicle type")
	}
}

// Function to refuel a vehicle (Car or Plane)
func RefuelVehicleWithoutInterface(vehicle interface{}) {
	switch vehicle.(type) {
	case Car:
		fmt.Println("Refueling the car")
	case Plane:
		fmt.Println("Refueling the plane")
	default:
		fmt.Println("Unknown vehicle type")
	}
}

func main() {
	// Create instances of Car and Plane
	car := Car{FuelLevel: 20.0}
	plane := Plane{FuelLevel: 10000.0}

	// Drive and refuel the Car without using interfaces
	fmt.Println("Driving and refueling the Car:")
	DriveVehicleWithoutInterface(car)
	RefuelVehicleWithoutInterface(car)

	fmt.Println("\nDriving and refueling the Plane:")
	// Drive and refuel the Plane without using interfaces
	DriveVehicleWithoutInterface(plane)
	RefuelVehicleWithoutInterface(plane)
}

/*

Driving and refueling the Car:
Driving the car
Refueling the car

Driving and refueling the Plane:
Flying the plane
Refueling the plane

*/

/*

How bad is this code 

The DriveVehicleWithoutInterface and RefuelVehicleWithoutInterface functions directly handle concrete types (Car and Plane) based on the provided string identifier for the vehicle type.
Introducing a new vehicle type involves manually adjusting the function signature and body to include the new type.
Keep in mind that while this approach may be simpler in some cases, it lacks the benefits of abstraction and formal contracts provided by interfaces. Using interfaces typically provides a more organized and extensible solution.

*/

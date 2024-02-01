package main

import "fmt"

// Vehicle interface with separate methods for driving and refueling
type Vehicle interface {
	Drive()
	Refuel()
}

// Car type implementing the Vehicle interface
type Car struct {
	Model string
}

func (c Car) Drive() {
	fmt.Println("Driving", c.Model)
}

func (c Car) Refuel() {
	fmt.Println("Refueling", c.Model)
}

// ElectricCar type implementing the Vehicle interface with a different refueling behavior
type ElectricCar struct {
	Model string
}

func (ec ElectricCar) Drive() {
	fmt.Println("Driving", ec.Model)
}

func (ec ElectricCar) Refuel() {
	fmt.Println("Charging", ec.Model)
}

// Function demonstrating SRP and extensibility
func DriveAndRefuel(v Vehicle) {
	v.Drive()
	v.Refuel()
}

func main() {
	// Original Car
	myCar := Car{"XYZ"}
	DriveAndRefuel(myCar)

	// New ElectricCar with different refueling behavior
	myElectricCar := ElectricCar{"ABC"}
	DriveAndRefuel(myElectricCar)
}

/*

Driving XYZ
Refueling XYZ
Driving ABC
Charging ABC


*/


/*

Explain this example

We have the Vehicle interface with separate Drive and Refuel methods.
The Car struct implements the Vehicle interface with its specific driving and refueling behavior.
We introduce a new type, ElectricCar, which also implements the Vehicle interface. This type has a different refueling behavior (charging).
The DriveAndRefuel function takes any type that satisfies the Vehicle interface and performs driving and refueling without knowing the specific type.
With this design, you can easily introduce new types of vehicles (e.g., ElectricCar, HybridCar) without modifying the existing DriveAndRefuel function or the Car implementation. Each new type can have its own behavior for driving and refueling, demonstrating adherence to the Single Responsibility Principle and allowing for easy extension without modifying existing code.

*/



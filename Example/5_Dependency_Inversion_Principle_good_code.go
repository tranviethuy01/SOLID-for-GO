package main

import "fmt"

// Switchable interface with TurnOn and TurnOff methods
type Switchable interface {
	TurnOn()
	TurnOff()
}

// LightSwitch type implementing the Switchable interface
type LightSwitch struct {
	State bool
}

func (ls *LightSwitch) TurnOn() {
	ls.State = true
	fmt.Println("Light is ON")
}

func (ls *LightSwitch) TurnOff() {
	ls.State = false
	fmt.Println("Light is OFF")
}

// Fan type implementing the Switchable interface
type Fan struct {
	State bool
}

func (f *Fan) TurnOn() {
	f.State = true
	fmt.Println("Fan is ON")
}

func (f *Fan) TurnOff() {
	f.State = false
	fmt.Println("Fan is OFF")
}

// Function demonstrating the Dependency Inversion Principle

func ControlSwitch (s Switchable, on bool) {
	if on {
		s.TurnOn()
	} else {
		s.TurnOff()
	}
}


func main() {
	// Create instances of LightSwitch and Fan
	lightSwitch := &LightSwitch{}
	fan := &Fan{}

	// Demonstrate using the ControlSwitch function with different types
	fmt.Println("Turning ON the LightSwitch:")
	ControlSwitch(lightSwitch, true)

	fmt.Println("\nTurning OFF the Fan:")
	ControlSwitch(fan, false)
}



/*


Turning ON the LightSwitch:
Light is ON

Turning OFF the Fan:
Fan is OFF

*/


/*

Explain this example:

We have the Switchable interface with TurnOn and TurnOff methods.
The LightSwitch and Fan types implement the Switchable interface, providing their own implementations for turning on and off.
The ControlSwitch function depends on the Switchable interface instead of concrete types. It can work with any type that implements the Switchable interface without knowing the specific type.
The main function demonstrates using the ControlSwitch function with instances of both LightSwitch and Fan without modifying the function.
By depending on the Switchable interface, the ControlSwitch function adheres to the Dependency Inversion Principle (DIP). This allows you to introduce new types implementing Switchable without modifying the existing code, promoting flexibility and extensibility.



*/

package main

import "fmt"

// LightSwitch type
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

// Fan type
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

// Function without using an interface
func ControlSwitchWithoutInterface(switch interface{}, on bool) {
	switchable, ok := switch.(SwitchableWithoutInterface)
	if !ok {
		fmt.Println("Invalid switch type")
		return
	}

	if on {
		switchable.TurnOn()
	} else {
		switchable.TurnOff()
	}
}

type SwitchableWithoutInterface interface {
	TurnOn()
	TurnOff()
}

func main() {
	// Create instances of LightSwitch and Fan
	lightSwitch := &LightSwitch{}
	fan := &Fan{}

	// Demonstrate using the ControlSwitchWithoutInterface function with different types
	fmt.Println("Turning ON the LightSwitch:")
	ControlSwitchWithoutInterface(lightSwitch, true)

	fmt.Println("\nTurning OFF the Fan:")
	ControlSwitchWithoutInterface(fan, false)
}


/*

How bad this code 

There is no explicit Switchable interface.
The ControlSwitchWithoutInterface function uses an empty interface (interface{}) to accept any type, and it performs a type assertion to check if the type implements SwitchableWithoutInterface.
The SwitchableWithoutInterface type is introduced to explicitly declare the contract for types that can be controlled using the ControlSwitchWithoutInterface function.
While this approach allows you to work with different types, it lacks the clear contract and benefits of interfaces. The type assertion checks and the introduction of SwitchableWithoutInterface represent an attempt to provide a form of contract, but it's less explicit and can lead to more runtime errors. Interfaces provide a more formal and compile-time-checked way to define contracts between components.

*/

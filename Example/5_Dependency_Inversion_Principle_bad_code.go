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

// Function without using any interface
func ControlSwitchWithoutInterface(switchType string, on bool) {
	switch switchType {
	case "LightSwitch":
		lightSwitch := &LightSwitch{}
		if on {
			lightSwitch.TurnOn()
		} else {
			lightSwitch.TurnOff()
		}
	case "Fan":
		fan := &Fan{}
		if on {
			fan.TurnOn()
		} else {
			fan.TurnOff()
		}
	default:
		fmt.Println("Unknown switch type")
	}
}

func main() {
	// Demonstrate using the ControlSwitchWithoutInterface function with different types
	fmt.Println("Turning ON the LightSwitch:")
	ControlSwitchWithoutInterface("LightSwitch", true)

	fmt.Println("\nTurning OFF the Fan:")
	ControlSwitchWithoutInterface("Fan", false)
}
/*

Turning ON the LightSwitch:
Light is ON

Turning OFF the Fan:
Fan is OFF

*/


/*
How bad this example:

The ControlSwitchWithoutInterface function directly creates instances of LightSwitch and Fan based on the provided string identifier for the switch type.
The function then calls the appropriate methods based on the switch type.
While this approach avoids interfaces entirely, it has some downsides. Adding new switch types involves modifying the ControlSwitchWithoutInterface function, which violates the Open/Closed Principle (OCP). Interfaces provide a more extensible and modular way to achieve the same goal.



*/

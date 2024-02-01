package main

import "fmt"

// Workable interface with a Work method
type Workable interface {
	Work()
}

// Eatable interface with an Eat method
type Eatable interface {
	Eat()
}

// Worker type implementing both Workable and Eatable interfaces
type Worker struct{}

func (w Worker) Work() {
	fmt.Println("Working")
}

func (w Worker) Eat() {
	fmt.Println("Eating during break")
}

// Robot type implementing only the Workable interface
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Performing work tasks")
}

// Function demonstrating the use of interfaces
func PerformActivities(activity interface{}) {
	switch a := activity.(type) {
	case Workable:
		a.Work()
	case Eatable:
		a.Eat()
	default:
		fmt.Println("Unknown activity")
	}
}

func main() {
	// Create instances of Worker and Robot
	worker := Worker{}
	robot := Robot{}

	// Demonstrate using interfaces
	fmt.Println("Worker's activities:")
	PerformActivities(worker)

	fmt.Println("\nRobot's activities:")
	PerformActivities(robot)
}

/*


Worker's activities:
Working
=> fail cmnr


Robot's activities:
Performing work tasks


*/

/*
Explain this example:

We have two interfaces, Workable and Eatable, each with a single method.
The Worker type implements both interfaces, providing specific implementations for Work and Eat methods.
The Robot type only implements the Workable interface, as robots may not eat.
The PerformActivities function takes an interface (activity) and uses a type switch to determine the specific interface implemented by the type and invokes the corresponding method.
This design adheres to the Interface Segregation Principle (ISP) by allowing types to implement only the methods they need. In this way, interfaces stay focused on specific responsibilities, promoting cleaner and more modular code.



*/

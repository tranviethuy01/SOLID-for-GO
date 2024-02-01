package main

import "fmt"

// Worker type with Work and Eat methods
type Worker struct{}

func (w Worker) Work() {
	fmt.Println("Working")
}

func (w Worker) Eat() {
	fmt.Println("Eating during break")
}

// Robot type with only a Work method
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Performing tasks")
}

// Function demonstrating without interfaces
func PerformActivitiesWithoutInterface(activity interface{}) {
	switch a := activity.(type) {
	case Worker:
		a.Work()
		a.Eat()
	case Robot:
		a.Work()
	default:
		fmt.Println("Unknown activity")
	}
}

func main() {
	// Create instances of Worker and Robot
	worker := Worker{}
	robot := Robot{}

	// Demonstrate without interfaces
	fmt.Println("Worker's activities:")
	PerformActivitiesWithoutInterface(worker)

	fmt.Println("\nRobot's activities:")
	PerformActivitiesWithoutInterface(robot)
}

/*

Worker's activities:
Working
Eating during break

Robot's activities:
Performing tasks

*/


/*

Explain this example:

The Worker type has both Work and Eat methods.
The Robot type only has a Work method.
The PerformActivitiesWithoutInterface function uses a type switch to handle both Worker and Robot types directly.
This approach works without formal interfaces, but it lacks the clarity and separation of concerns that interfaces provide. In the absence of interfaces, you need to check the type explicitly within the switch statement and handle each case individually. As the number of types and methods grows, this can lead to less maintainable and more error-prone code. Interfaces offer a more organized and extensible way to handle varying behavior for different types.

*/



